import{CommentService} from './commentservice.js'
import{expect,test} from 'vitest'
test("listComments",async()=>{
    const auth=new CommentService()
    const postID='56'
     expect(!await auth.listCommentsByPost(postID))
})