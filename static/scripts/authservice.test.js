import{expect,test} from 'vitest'
import {AuthService} from './authservice.js'

test("login test",async()=>{
    const auth=new AuthService()
    const cred={email:'',password:'78'}
   const data = { error: true, message: "Please provide both email and password!"}
     expect(await auth.login(cred)).toStrictEqual( data)
})
test("login wrong creds",async()=>{
  const auth=new AuthService()
  const cred={email:'Mamapima@gmaail.com',password:'785757575757'}
 const data = "wrong email or passsword"
   expect(!await auth.login(cred))
})
test("register test",async()=>{
  const auth=new AuthService()
  const cred={email:'',password:'78',user_name:"rector"}
 const data = { error: true, message: "Please provide all required fields!"}
   expect(!await auth.login(cred))
})


