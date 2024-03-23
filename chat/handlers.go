package chat

import (
    "fmt"
    "net/http"
    "html/template"

    "go_authentication/options"
    "go_authentication/authtoken"
)


func UsChat(w http.ResponseWriter, r *http.Request) {

    cls,err := authtoken.OnToken(w,r)
    if cls == nil {
        return
    }
    if err != nil {
        return
    }
    id,err := options.IdUrl(w,r)
    if err != nil {
        return
    }

    // detail..
    user,completed,err := qrUsSscCh(w,id)
    if err != nil {
        return
    }
    in_user,err := idUs(w,user,completed)
    if err != nil {
        return
    }
    // ..

    rows,err := qUsCh(w,cls.User_id,id)
    if err != nil {
        return
    }
    names,err := usChat(w,rows)
    if err != nil {
        return
    }

    type ListData struct {
        Ssc int
        I []*MsgUser
        Uid int
        Uemail string
        T []*ChUser
    }
    data := ListData {
        Ssc: id,
        I: names,
        Uid: cls.User_id,
        Uemail: cls.Email,
        T: in_user,
    }

    if r.Method == "GET" {
        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/chat/user.html", "./tpl/base.html" ))

        tpl.ExecuteTemplate(w, "base", data)
    }
}

func GrChat(w http.ResponseWriter, r *http.Request) {

    cls,err := authtoken.OnToken(w,r)
    if cls == nil {
        return
    }
    if err != nil {
        return
    }

    id,err := options.IdUrl(w,r)
    if err != nil {
        return
    }

    type ListData struct {
        I []*MsgGroup
        D Group
    }

    idname,err := idGroup(w,id)
    if err != nil {
        return
    }

    owner := cls.User_id

    rows,err := qGrCh(w,owner,id)
    if err != nil {
        return
    }
    names,err := groupChat(w,rows,owner,id)
    if err != nil {
        return
    }

    data := ListData {
        I: names,
        D: idname,
    }

    if r.Method == "GET" {
        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/chat/group.html", "./tpl/base.html" ))

        tpl.ExecuteTemplate(w, "base", data)
    }
}


func GrAll(w http.ResponseWriter, r *http.Request) {

    if r.Method == "GET" {

        cls,err := authtoken.OnToken(w,r)

        if cls != nil && err == nil {

        rows,err := qGroup(w)
        if err != nil {
            return
        }
        names,err := allGroup(w,rows)
        if err != nil {
            return
        }
        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/chat/all_group.html", "./tpl/base.html" ))
        tpl.ExecuteTemplate(w, "base", names)

        }
    }
}
func GrOwr(w http.ResponseWriter, r *http.Request) {

    if r.Method == "GET" {

        cls,err := authtoken.ListToken(w,r)

        if cls != nil && err == nil {

        owner := cls.User_id

        rows,err := qUsGroup(w,owner)
        if err != nil {
            return
        }
        names,err := userGroup(w,rows)
        if err != nil {
            return
        }
        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/chat/all_owr_group.html", "./tpl/base.html" ))
        tpl.ExecuteTemplate(w, "base", names)

        } else {

        rows,err := qGroup(w)
        if err != nil {
            return
        }
        names,err := allGroup(w,rows)
        if err != nil {
            return
        }
        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/chat/all.html", "./tpl/base.html" ))
        tpl.ExecuteTemplate(w, "base", names)

        }
    }
}


func DtlGr(w http.ResponseWriter, r *http.Request) {

    if r.Method == "GET" {

        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/chat/detail.html", "./tpl/base.html" ))

        id,err := options.IdUrl(w,r)
        if err != nil {
            return
        }
        
        i,err := idGroup(w,id)
        if err != nil {
            return
        }

        cls := authtoken.WhoisWho(w,r)
        
        type ListData struct {
            Auth string
            I Group
        }
        if cls != nil {
            data := ListData {
                Auth: cls.Email,
                I: i,
            }
            fmt.Println("is cls")
            tpl.ExecuteTemplate(w, "base", data)
        } else {
            data := ListData {
                I: i,
            }
            fmt.Println("ne cls", cls)
            tpl.ExecuteTemplate(w, "base", data)
        }
    }
}


func HomeChat(w http.ResponseWriter, r *http.Request) {

    if r.Method == "GET" {
        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/chat/index.html", "./tpl/base.html" ))

        tpl.ExecuteTemplate(w, "base", nil)
    }
}