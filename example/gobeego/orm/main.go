package main

import (
	"github.com/astaxie/beego/orm"
        "fmt"
)

func testOneOneCRUD(){
    
    o := orm.NewOrm()

    o.Raw("delete from \"user\"").Exec()
    o.Raw("delete from profile").Exec()
    
    profile := new(Profile)
    profile.Age = 31
    profile.Money = 4321.09
    o.Insert(profile)
    fmt.Println(*profile)

    user := new(User)
    user.UserName = "slene"
    user.Email = "vslene@gmail.com"
    user.Password = "pass"
    user.Status = 1
    user.IsStaff = false
    user.IsActive = true
    user.Profile = profile

    id, err := o.Insert(user)
    if err != nil{
 	fmt.Println(err)
	return
    }

    user1 := User{ID: int(id)}
    o.Read(&user1)
    fmt.Println(user1)
    if user1.Profile != nil{
        fmt.Println(user1.Profile.Age)
        o.Read(user1.Profile)
        fmt.Println(user1.Profile.Age)
    }

    //o.Delete(user)
    o.Delete(profile)
}


func testOneManyCRUD(){
    o := orm.NewOrm()
    o.Raw(`delete from "user"`).Exec()
    o.Raw("delete from post").Exec()

    user := new(User)
    user.UserName = "slene"
    user.Email = "vslene@gmail.com"
    user.Password = "pass"
    user.Status = 1
    user.IsStaff = false
    user.IsActive = true
    o.Insert(user)
    posts := []*Post{
                {User: user, Tags: []*Tag{}, Title: "Introduction", Content: "Introduction content"},
                {User: user, Tags: []*Tag{}, Title: "Examples", Content: `Exampe content`},
        }
    o.InsertMulti(100, posts) 
    
    o.Delete(user)
}


func testQuery(){
    //RelatedSel
    postTag := PostTags{ID:1}
    o := orm.NewOrm()
    qs := o.QueryTable(postTag)
    //SELECT T0."id", T0."post_id", T0."tag_id", T1."id", T1."user_id", T1."title", T1."content", T1."created", T1."updated", T2."id", T2."user_name", T2."email", T2."password", T2."Status", T2."is_staff", T2."is_active", T2."created", T2."updated", T2."profile_id", T2."nums", T3."id", T3."name", T3."best_post_id", T4."id", T4."user_id", T4."title", T4."content", T4."created", T4."updated" FROM "post_tags" T0 INNER JOIN "post" T1 ON T1."id" = T0."post_id" INNER JOIN "user" T2 ON T2."id" = T1."user_id" INNER JOIN "tag" T3 ON T3."id" = T0."tag_id" LEFT OUTER JOIN "post" T4 ON T4."id" = T3."best_post_id" LIMIT 1
    qs.RelatedSel().One(&postTag)
    //SELECT T0."id", T0."post_id", T0."tag_id", T1."id", T1."user_id", T1."title", T1."content", T1."created", T1."updated" FROM "post_tags" T0 INNER JOIN "post" T1 ON T1."id" = T0."post_id" LIMIT 1
    qs.RelatedSel("post").One(&postTag)       
    //SELECT T0."id", T0."post_id", T0."tag_id", T1."id", T1."user_id", T1."title", T1."content", T1."created", T1."updated", T2."id", T2."user_name", T2."email", T2."password", T2."Status", T2."is_staff", T2."is_active", T2."created", T2."updated", T2."profile_id", T2."nums" FROM "post_tags" T0 INNER JOIN "post" T1 ON T1."id" = T0."post_id" INNER JOIN "user" T2 ON T2."id" = T1."user_id" LIMIT 1
    qs.RelatedSel("post", "post__user").One(&postTag)       
    //SELECT T0."id", T0."post_id", T0."tag_id", T1."id", T1."user_id", T1."title", T1."content", T1."created", T1."updated", T2."id", T2."name", T2."best_post_id" FROM "post_tags" T0 INNER JOIN "post" T1 ON T1."id" = T0."post_id" INNER JOIN "tag" T2 ON T2."id" = T0."tag_id" LIMIT 1
    qs.RelatedSel("post", "tag").One(&postTag)       
  
    user := new(User)
    user.UserName = "slene"
    user.Email = "vslene@gmail.com"
    user.Password = "pass"
    user.Status = 1
    user.IsStaff = false
    user.IsActive = true
    o.Insert(user) 
    var users []*User
    _, _ = o.QueryTable("user").All(&users)
    fmt.Println(users) 
    var users1 []User
    o.QueryTable("user").All(&users1)
    fmt.Println(users1) 
    o.Raw(`delete from "user"`).Exec()
}

func main(){
    //testOneOneCRUD()
    //testOneManyCRUD()
    testQuery()
}
