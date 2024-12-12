export namespace statistics {
	
	export class OutcomeAggregator {
	    Count: number;
	    FailureCount: number;
	    Sum: number;
	    SumOfSquares: number;
	    PrecisionTarget: number;
	    WindowSize: number;
	    MeanWindow: number[];
	    // Go type: sync
	    Mu: any;
	    ppf: number;
	    mean: number;
	    variance: number;
	    min: number;
	    q1: number;
	    q2: number;
	    q3: number;
	    max: number;
	    LearningRate: number;
	    stable: boolean;
	
	    static createFrom(source: any = {}) {
	        return new OutcomeAggregator(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Count = source["Count"];
	        this.FailureCount = source["FailureCount"];
	        this.Sum = source["Sum"];
	        this.SumOfSquares = source["SumOfSquares"];
	        this.PrecisionTarget = source["PrecisionTarget"];
	        this.WindowSize = source["WindowSize"];
	        this.MeanWindow = source["MeanWindow"];
	        this.Mu = this.convertValues(source["Mu"], null);
	        this.ppf = source["ppf"];
	        this.mean = source["mean"];
	        this.variance = source["variance"];
	        this.min = source["min"];
	        this.q1 = source["q1"];
	        this.q2 = source["q2"];
	        this.q3 = source["q3"];
	        this.max = source["max"];
	        this.LearningRate = source["LearningRate"];
	        this.stable = source["stable"];
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

