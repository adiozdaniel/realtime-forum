import {API_ENDPOINTS, userData} from "./data.js";

// AuthService class for handling authentication-related API requests
class AuthService {
	constructor() {
		this.apiEndpoints = API_ENDPOINTS;
	}
}

// Method to log in a user
AuthService.prototype.login = async function (credentials) {
	if (!credentials?.email || !credentials?.password) {
		return (data = {
			error: true,
			message: "Please provide both email and password!",
		});
	}

	try {
		const response = await fetch(this.apiEndpoints.login, {
			method: "POST",
			headers: {
				"Content-Type": "application/json",
			},
			body: JSON.stringify(credentials),
		});

		const res = response.json();
		userData = res;

		return res;
	} catch (error) {
		return error;
	}
};

// Method to register a new user
AuthService.prototype.register = async function (userData) {
	if (!userData?.email || !userData?.password || !userData?.user_name) {
		return (data = {
			error: true,
			message: "Please provide all required fields!",
		});
	}

	try {
		const response = await fetch(this.apiEndpoints.register, {
			method: "POST",
			body: JSON.stringify(userData),
			headers: {
				"Content-Type": "application/json",
			},
		});

		const res = response.json();
		userData = res;

		return res;
	} catch (error) {
		return error;
	}
};

// Method to log out a user
AuthService.prototype.logout = async function () {
	try {
		const response = await fetch(this.apiEndpoints.logout, {
			method: "POST",
		});

		return response.json();
	} catch (error) {
		return error;
	}
};

// Method to check if the user is authenticated
AuthService.prototype.isAuthenticated = async function () {
	try {
		const response = await fetch(this.apiEndpoints.check, {
			method: "GET",
		});

		return response.json();
	} catch (error) {
		return error;
	}
};

export { AuthService };
