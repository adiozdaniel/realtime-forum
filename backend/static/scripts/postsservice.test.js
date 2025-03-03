import{test,expect} from 'vitest'
import { PostService } from './postsservice'
test("fetchPOsts",async()=>{
  const post=new PostService
  expect(await post.fetchPosts())  
})