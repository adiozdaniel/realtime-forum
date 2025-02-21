import{expect,test} from 'vitest'
import {AuthService} from './authservice.js'

test("login test",async()=>{
    const auth=new AuthService()
    const cred={email:'',password:'78'}
   const data = { error: true, message: "Please provide both email and password!"}
     expect(await auth.login(cred)).toStrictEqual( data)
})

