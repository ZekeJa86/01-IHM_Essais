export namespace main {
	
	export class ContactForm {
	    firstName: string;
	    lastName: string;
	    company: string;
	    email: string;
	    country: string;
	    phoneNumber: string;
	    message: string;
	    agreeToPolicy: boolean;
	
	    static createFrom(source: any = {}) {
	        return new ContactForm(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.firstName = source["firstName"];
	        this.lastName = source["lastName"];
	        this.company = source["company"];
	        this.email = source["email"];
	        this.country = source["country"];
	        this.phoneNumber = source["phoneNumber"];
	        this.message = source["message"];
	        this.agreeToPolicy = source["agreeToPolicy"];
	    }
	}
	export class EmailResponse {
	    success: boolean;
	    message: string;
	
	    static createFrom(source: any = {}) {
	        return new EmailResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.success = source["success"];
	        this.message = source["message"];
	    }
	}

}

