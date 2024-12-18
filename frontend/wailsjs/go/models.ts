export namespace statistics {
	
	export class LearnedSummary {
	    Stable: boolean;
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

