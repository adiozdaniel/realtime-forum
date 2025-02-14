import { AuthService } from "./authservice.js"

class Authmiddleware {
    constructor(){
        this.authService = new AuthService();
        this.authButton = document.querySelector(".sign-in-button");
    }
}

//Authchecker method to check if user is authenticated
Authmiddleware.prototype.authChecker = async function(){
    const isAuthenticated = await this.authService.isAuthenticated();
    
    if (isAuthenticated.error) console.log(isAuthenticated.message);

    if (isAuthenticated?.data) {
        localStorage.setItem('userdata', isAuthenticated.data);
        this.authButton.textContent = isAuthenticated.data ? "Sign Out" : "Sign In";
    };
}

const authmiddleware = new Authmiddleware();

await authmiddleware.authChecker();