import { AuthService } from "./authservice.js";

class Authmiddleware {
	constructor() {
		this.authService = new AuthService();
	}

	// Authchecker method to check if the user is authenticated
	async authChecker() {
		const isAuthenticated = await this.authService.isAuthenticated();

		if (isAuthenticated.error) {
			console.log(isAuthenticated.message);
			return null;
		}

		if (isAuthenticated?.data) {
			return isAuthenticated.data;
		}

		return null;
	}
}

const authmiddleware = new Authmiddleware();

// Export a function to get user data instead of a resolved variable
async function getUserData() {
	return await authmiddleware.authChecker();
}

export { getUserData };
