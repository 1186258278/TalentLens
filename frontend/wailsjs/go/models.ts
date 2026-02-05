export namespace main {
	
	export class AIConfig {
	    provider: string;
	    base_url: string;
	    api_key: string;
	    model: string;
	    max_retries: number;
	    timeout: number;
	
	    static createFrom(source: any = {}) {
	        return new AIConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.provider = source["provider"];
	        this.base_url = source["base_url"];
	        this.api_key = source["api_key"];
	        this.model = source["model"];
	        this.max_retries = source["max_retries"];
	        this.timeout = source["timeout"];
	    }
	}
	export class AnalysisResult {
	    overall_score: number;
	    skill_match: number;
	    experience_match: number;
	    education_match: number;
	    recommendation: string;
	    strengths: string[];
	    weaknesses: string[];
	    summary: string;
	    analyzed_at: string;
	
	    static createFrom(source: any = {}) {
	        return new AnalysisResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.overall_score = source["overall_score"];
	        this.skill_match = source["skill_match"];
	        this.experience_match = source["experience_match"];
	        this.education_match = source["education_match"];
	        this.recommendation = source["recommendation"];
	        this.strengths = source["strengths"];
	        this.weaknesses = source["weaknesses"];
	        this.summary = source["summary"];
	        this.analyzed_at = source["analyzed_at"];
	    }
	}
	export class JobConfig {
	    title: string;
	    requirements: string[];
	    required_skills: string[];
	    experience_years: number;
	    education_level: string;
	
	    static createFrom(source: any = {}) {
	        return new JobConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.title = source["title"];
	        this.requirements = source["requirements"];
	        this.required_skills = source["required_skills"];
	        this.experience_years = source["experience_years"];
	        this.education_level = source["education_level"];
	    }
	}
	export class Config {
	    ai: AIConfig;
	    job: JobConfig;
	
	    static createFrom(source: any = {}) {
	        return new Config(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ai = this.convertValues(source["ai"], AIConfig);
	        this.job = this.convertValues(source["job"], JobConfig);
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
	
	export class Resume {
	    id: string;
	    file_name: string;
	    file_path: string;
	    file_type: string;
	    file_size: number;
	    content: string;
	    status: string;
	    score: number;
	    analysis?: AnalysisResult;
	    // Go type: time
	    created_at: any;
	
	    static createFrom(source: any = {}) {
	        return new Resume(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.file_name = source["file_name"];
	        this.file_path = source["file_path"];
	        this.file_type = source["file_type"];
	        this.file_size = source["file_size"];
	        this.content = source["content"];
	        this.status = source["status"];
	        this.score = source["score"];
	        this.analysis = this.convertValues(source["analysis"], AnalysisResult);
	        this.created_at = this.convertValues(source["created_at"], null);
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

