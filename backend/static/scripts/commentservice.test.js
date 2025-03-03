import{CommentService} from './commentservice.js'
import{expect,test} from 'vitest'
test("listComments",async()=>{
    const auth=new CommentService()
    const postID='56'
     expect(!await auth.listCommentsByPost(postID))
})
test("updatecomment",async()=>{
    const auth=new CommentService()
    const message={
        error: true,
        message: "Failed to edit comment. Please try again.",
    }
     expect(await auth.updateComment({})).toEqual(message)
})
test("deletecomment",async()=>{
    const auth=new CommentService()
    const message={
        error: true,
        message: "Failed to delete comment. Please try again.",
    }
     expect(await auth.deleteComment({})).toEqual(message)
})