export namespace simulation {
	
	export class AccountResults {
	    Name: string;
	    InvestmentResults: {[key: string]: statistics.LearnedSummary};
	
	    static createFrom(source: any = {}) {
	        return new AccountResults(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Name = source["Name"];
	        this.InvestmentResults = this.convertValues(source["InvestmentResults"], statistics.LearnedSummary, true);
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
	export class SimulationResult {
	    YearlyResults: AccountResults[];
	    TotalSims: number;
	    SimulationDuration: number;
	
	    static createFrom(source: any = {}) {
	        return new SimulationResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.YearlyResults = this.convertValues(source["YearlyResults"], AccountResults);
	        this.TotalSims = source["TotalSims"];
	        this.SimulationDuration = source["SimulationDuration"];
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

export namespace statistics {
	
	export class LearnedSummary {
	    Stable: boolean;
	    Stability: number;
	    Confidence: number;
	    Count: number;
	    PPF: number;
	    Mean: number;
	    Variance: number;
	    Kurtosis: number;
	    Skewness: number;
	    Min: number;
	    Q1: number;
	    Q2: number;
	    Q3: number;
	    Max: number;
	
	    static createFrom(source: any = {}) {
	        return new LearnedSummary(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Stable = source["Stable"];
	        this.Stability = source["Stability"];
	        this.Confidence = source["Confidence"];
	        this.Count = source["Count"];
	        this.PPF = source["PPF"];
	        this.Mean = source["Mean"];
	        this.Variance = source["Variance"];
	        this.Kurtosis = source["Kurtosis"];
	        this.Skewness = source["Skewness"];
	        this.Min = source["Min"];
	        this.Q1 = source["Q1"];
	        this.Q2 = source["Q2"];
	        this.Q3 = source["Q3"];
	        this.Max = source["Max"];
	    }
	}

}

