export namespace main {
	
	export enum Weekday {
	    FRIDAY = "Friday",
	    MONDAY = "Monday",
	    SATURDAY = "Saturday",
	    SUNDAY = "Sunday",
	    THURSDAY = "Thursday",
	    TUESDAY = "Tuesday",
	    WEDNESDAY = "Wednesday",
	}
	export class Address {
	    city: string;
	    country: string;
	
	    static createFrom(source: any = {}) {
	        return new Address(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.city = source["city"];
	        this.country = source["country"];
	    }
	}
	export class Person {
	    name: string;
	    age: number;
	    address?: Address;
	    numTimesLoggedIn: number;
	
	    static createFrom(source: any = {}) {
	        return new Person(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.age = source["age"];
	        this.address = this.convertValues(source["address"], Address);
	        this.numTimesLoggedIn = source["numTimesLoggedIn"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

