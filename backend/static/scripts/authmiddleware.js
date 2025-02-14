import { AuthService } from "./authservice.js"
import { userData } from "./data.js";

class Authmiddleware {
    constructor(){
        this.authService = new AuthService();
    }
}

//Authchecker method to check if user is authenticated
Authmiddleware.prototype.authChecker = async function(){
    const isAuthenticated = await this.authService.isAuthenticated();
    
    if (isAuthenticated.error) console.log(isAuthenticated.message);
    if (isAuthenticated?.data) userData = isAuthenticated.data;
}

document.addEventListener("DOMContentLoaded", async() => {
    const authmiddleware = new Authmiddleware();

    await authmiddleware.authChecker();
});
