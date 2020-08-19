package model

import (
	"SNSProject/DB"
	"fmt"
	b "github.com/orca-zhang/borm"
	"strings"
)

//帖子
type Post struct {
	BormLastId int16
	Title string `borm:"title"`
	Pid int16 `borm:"pid"`
	Uperid int16 `borm:"uperid"`
	Class string `borm:"class"`
	Comment int `borm:"comment"`
	Content string `borm:"content"`
	Files string `borm:"files"`
	IsLike bool
}

//帖子用户链接 点赞链接也能用
type PULink struct {
	Pid int16 `borm:"pid"`
	Uid int16 `borm:"uid"`
}

//帖子分类
type PClass struct {
	Class string `borm:"class"`
	Cid int `borm:"cid"`
}

type Comment struct {
	Content string `borm:"content"`
	Cid int16 `borm:"cid"`
	Uid int16 `borm:"uid"`
	Pid int16 `borm:"pid"`
}

//插入帖子
func PostUp(post Post) error {
	table := b.Table(DB.DB, "post").Debug()

	_, err := table.Insert(&post)

	if err != nil {
		fmt.Print(err)
		return err
	}

	err = PULinker(post.BormLastId, post.Uperid)
	if err != nil {
		fmt.Print(err)
		return err
	}

	return err
}

//post user 关联
func PULinker(pid int16, uid int16) error {
	pulin := PULink{
		Pid:pid,
		Uid:uid,
	}

	table := b.Table(DB.DB, "post_user_link")
	_, err := table.Insert(&pulin)
	if err != nil {
		fmt.Print(err)
	}
	return err
}

//查询post
func QueryPost(pid int16) (Post, error) {
	table := b.Table(DB.DB, "post").Debug()

	var posts []Post
	var post Post
	count, err := table.Select(&posts, b.Where("pid = ?", pid))

	if err != nil {
		fmt.Print(err)
		return post, err
	}

	if count <= 0 {
		return post, err
	}

	post = posts[0]
	return post, err
}

//帖子点赞
func SetPostCommend(add int, uid int16, pid int16) error {
	table := b.Table(DB.DB, "post").Debug()

	//var uStr b.U
	//if add == 1 {
	//	uStr = b.U("comment+1")
	//}else if add == 0 {
	//	uStr = b.U("comment-1")
	//}
	// 使用map更新
	//_, err := table.Update(b.V{
	//	"comment":  uStr, // 使用b.U来处理非变量更新
	//}, b.Where(b.Eq("pid", pid)))
	//if err != nil {
	//	return err
	//}
	_, err := QueryPost(pid)
	if err != nil {
		return err
	}

	var count int
	count, err = LinkPostLikes(add, pid, uid)
	if err != nil {
		return err
	}

	_, err = table.Update(b.V{
		"comment": count,
	}, b.Where(b.Eq("pid",pid)))

	return err
}

//帖子点赞关联表更新
func LinkPostLikes(add int, pid int16, uid int16) (count int, err error) {
	table := b.Table(DB.DB, "post_likes").Debug()

	var liner = PULink{
		Uid:uid,
		Pid:pid,
	}

	if add == 1 {
		count, err = table.Insert(&liner)
	}else {
		count, err = table.Delete(b.Where(b.Eq("uid", liner.Uid), b.Eq("pid", liner.Pid)))
	}
	return count, err
}

//查询点赞列表
func QueryPostLikes(pid string, uid string) ([]PULink, error) {
	var likes []PULink

	table := b.Table(DB.DB, "post_likes").Debug()

	var err error
	if uid == "" {
		_, err = table.Select(&likes, b.Where("pid = ?", pid))
	}else {
		_, err = table.Select(&likes, b.Where("pid = ? & uid = ?", pid, uid))
	}

	return likes, err
}

//查询帖子列表
func QueryPostList(index int, size int, uid string) []Post {
	table := b.Table(DB.DB, "post").Debug()

	var posts []Post
	var count int = (index - 1) * size
	_, err := table.Select(&posts, b.Limit(count,size))

	if err != nil {
		fmt.Print(err)
	}

	var pids = make([]string, 0)
	for _, post := range posts {
		pids = append(pids, string(post.Pid))
	}

	var likes []PULink
	likes, err = QueryPostLikes(strings.Join(pids, ","), uid)

	for idx, post := range posts {
		for _, like := range likes {
			if post.Pid == like.Pid {
				posts[idx].IsLike = true
			}
		}
	}

	return posts;
}

//查询帖子分类
func QueryPostClass() []PClass {
	table := b.Table(DB.DB, "post_class")

	var pclass []PClass
	_, err := table.Select(&pclass, b.Where("cid>=100"))
	if err != nil {
		fmt.Print(err)
	}

	return pclass
}

//评论
func AddComment(commend Comment) error {
	table := b.Table(DB.DB, "comments")

	_, err := table.Insert(&commend)

	return err
}

func QueryCommentWithPid(pid int16, index, size int) (comments []Comment, err error) {
	table := b.Table(DB.DB, "comments")

	var count int = (index - 1) * size
	_, err = table.Select(&comments, b.Where("pid = ?", pid), b.Limit(count, size))

	return comments, err
}