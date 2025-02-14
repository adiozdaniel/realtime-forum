import { AuthService } from "./authservice.js";
import { CONSTANTS } from "./data.js";

class AuthHandler {

	constructor(){
    this.authService = new AuthService();

    // DOM Elements
    this.form = document.getElementById("signinForm");
    this.emailInput = document.getElementById("email");
    this.passwordInput = document.getElementById("password");
    this.passwordToggle = document.getElementById("passwordToggle");
    this.submitButton = this.form.querySelector('button[type="submit"]');
    this.spinner = this.submitButton.querySelector(".spinner");
	
    // Bind methods to preserve 'this' in event handlers
    this.togglePasswordVisibility = this.togglePasswordVisibility.bind(this);
    this.handleLogin = this.handleLogin.bind(this);
    this.clearErrors = this.clearErrors.bind(this);
 };	
}

// Initialize event listeners
AuthHandler.prototype.init = function () {
    this.passwordToggle.addEventListener("click", this.togglePasswordVisibility);
    this.form.addEventListener("submit", this.handleLogin);
    this.form.addEventListener("input", this.clearErrors);
};

// Toggle password visibility
AuthHandler.prototype.togglePasswordVisibility = function () {
    const type = this.passwordInput.type === "password" ? "text" : "password";
    this.passwordInput.type = type;

    const icon = this.passwordToggle.querySelector("i");
    if (icon) {
        icon.setAttribute("data-lucide", type === "password" ? "eye" : "eye-off");
    }
};

// Validate email format
AuthHandler.prototype.validateEmail = function (email) {
    return /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email);
};

// Display error message next to input fields
AuthHandler.prototype.showError = function (input, message) {
    const formGroup = input.closest(".form-group");
    let errorDiv = formGroup.querySelector(".error");

    if (!errorDiv) {
        errorDiv = document.createElement("div");
        errorDiv.className = "error";
        formGroup.appendChild(errorDiv);
    }

    errorDiv.textContent = message;
    input.classList.add("error-input");
};

// Remove error message when user starts typing
AuthHandler.prototype.removeError = function (input) {
    const formGroup = input.closest(".form-group");
    const error = formGroup.querySelector(".error");
    if (error) error.remove();
    input.classList.remove("error-input");
};

// Clear errors dynamically when user types
AuthHandler.prototype.clearErrors = function (event) {
    if (event.target.matches("#email, #password")) {
        this.removeError(event.target);
    }
};

// Handle form submission for user login
AuthHandler.prototype.handleLogin = async function (e) {
    e.preventDefault();
    this.removeError(this.emailInput);
    this.removeError(this.passwordInput);

    if (!this.validateEmail(this.emailInput.value)) {
        this.showError(this.emailInput, "Please enter a valid email address");
        return;
    }

    if (this.passwordInput.value.length < CONSTANTS.MIN_PASSWORD_LENGTH) {
        this.showError(this.passwordInput, `Password must be at least ${CONSTANTS.MIN_PASSWORD_LENGTH} characters`);
        return;
    }

    this.submitButton.disabled = true;
    this.spinner.classList.remove("hidden");
    this.submitButton.querySelector("span").textContent = "Signing in...";

    const formData = {
        email: this.emailInput.value,
        password: this.passwordInput.value,
    };

    const response = await this.authService.login(formData);

	if (response.error) {
		console.error("Login failed:", response.message || "Unknown error");
        this.showError(this.emailInput, response.message || "Invalid email or password");
	}

	if (response?.data) {
		console.log("Login successful:", response.message);
        window.location.href = "/";
	} else {
        console.error("Login failed:", response.message || "Unknown error");
        this.showError(this.emailInput, response.message || "oops something went wrong");
    }

    this.submitButton.disabled = false;
    this.spinner.classList.add("hidden");
    this.submitButton.querySelector("span").textContent = "Sign in";
};

// Initialize the authentication handler on DOM load
document.addEventListener("DOMContentLoaded", function () {
    const auth = new AuthHandler();
    auth.init();
});
