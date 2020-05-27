package main

import (
     "github.com/astaxie/beego/orm"
     _ "github.com/lib/pq"
     "time"
)

type User struct {
        ID           int    `orm:"column(id)"`
        UserName     string `orm:"size(30);unique"`
        Email        string `orm:"size(100)"`
        Password     string `orm:"size(100)"`
        Status       int16  `orm:"column(Status)"`
        IsStaff      bool
        IsActive     bool      `orm:"default(true)"`
        Created      time.Time `orm:"auto_now_add;type(date)"`
        Updated      time.Time `orm:"auto_now"`
        Profile      *Profile  `orm:"null;rel(one);on_delete(cascade)"`
        Posts        []*Post   `orm:"reverse(many);on_delete(cascade)" json:"-"`
        ShouldSkip   string    `orm:"-"`
        Nums         int
        unexport     bool             `orm:"-"`
        unexportBool bool
}


type Profile struct {
        ID       int `orm:"column(id)"`
        Age      int16
        Money    float64
        User     *User `orm:"reverse(one);on_delete(cascade)" json:"-"`
        BestPost *Post `orm:"rel(one);null"`
}

type Tag struct {
        ID       int     `orm:"column(id)"`
        Name     string  `orm:"size(30)"`
        BestPost *Post   `orm:"rel(one);null"`
        Posts    []*Post `orm:"reverse(many)" json:"-"`
}

type Post struct {
        ID      int       `orm:"column(id)"`
        User    *User     `orm:"rel(fk)"`
        Title   string    `orm:"size(60)"`
        Content string    `orm:"type(text)"`
        Created time.Time `orm:"auto_now_add"`
        Updated time.Time `orm:"auto_now"`
        Tags    []*Tag    `orm:"rel(m2m);rel_through(main.PostTags)"`
}

type PostTags struct {
        ID   int   `orm:"column(id)"`
        Post *Post `orm:"rel(fk)"`
        Tag  *Tag  `orm:"rel(fk)"`
}



func init(){
        orm.Debug = true
	 // set default database
        orm.RegisterDataBase("default", "postgres", "user=test password=test dbname=test host=127.0.0.1 port=5432 sslmode=disable", 30)
        
        // register model
        orm.RegisterModel(new(User), new(Profile), new(Post), new(Tag), new(PostTags))
       
        //clean previous test data         
        o := orm.NewOrm()	
	o.Raw("drop table \"user\"").Exec()
        o.Raw("drop table profile").Exec()
        o.Raw("drop table post").Exec()
	o.Raw("drop table tag").Exec()
        o.Raw("drop table post_tags").Exec()

        // create or alert table
        orm.RunSyncdb("default", false, true)

}
