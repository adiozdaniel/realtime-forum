import{test,expect} from 'vitest'
import { PostService } from './postsservice'
test("fetchPOsts",async()=>{
  const post=new PostService
  expect(await post.fetchPosts())  
})
test("deletePost",async()=>{
    const post=new PostService
    expect(await post.deletePost({we:'hff'}))  
  })
  test("likePost",async()=>{
    const post=new PostService
    const postdata={}
    const message={
        error: true,
        message: "You need to login to like the post!",
    }
    expect(await post.likePost(postdata)).toEqual(message) 
  })
  test("dislikePost",async()=>{
    const post=new PostService
    const postdata={}
    const message={
        error: true,
        message: "You need to login to dislike the post!",
    }
    expect(await post.dislikePost(postdata)).toEqual(message) 
  })