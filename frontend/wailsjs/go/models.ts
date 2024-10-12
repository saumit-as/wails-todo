export namespace main {
	
	export class Todo {
	    id: number;
	    name: string;
	    completed: number;
	
	    static createFrom(source: any = {}) {
	        return new Todo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.completed = source["completed"];
	    }
	}

}

