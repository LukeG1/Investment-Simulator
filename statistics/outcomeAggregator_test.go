package statistics

import (
	"fmt"
	"math"
	"testing"
)

// func floatsAreClose(a, b, epsilon float64) bool {
// 	return math.Abs(a-b) <= epsilon
// }

// Take in the stock market values, calculate sample statistics, sample, check that they are similar
func TestAddOutcome(t *testing.T) {
	sampler := GenerateKernelSampler(&[]float64{0.43811155152887893, -0.0829794661190966, -0.25123636363636365, -0.4383754889178619, -0.08642364532019696, 0.49982225433526023, -0.011885656970912803, 0.4674042105263158, 0.3194341027550261, -0.3533672875436554, 0.29282654028436017, -0.010975646879756443, -0.10672873194221515, -0.1277145557655955, 0.19173762945914843, 0.25061310133060394, 0.1903067694944301, 0.358210843373494, -0.08429147465437781, 0.052, 0.057045751633986834, 0.18303223684210526, 0.30805539011316263, 0.2367846304454234, 0.18150988641144306, -0.012082047421904465, 0.525633212414349, 0.3259733185102835, 0.07439511873350935, -0.1045736018855796, 0.43719954988747184, 0.12056457163557326, 0.00336535314743695, 0.2663771295818275, -0.08811460517120888, 0.22611927099841514, 0.16415455878432425, 0.12399242477876114, -0.0997095423563779, 0.23802966513133328, 0.10814862651601535, -0.08241371076449064, 0.03561144905496419, 0.14221150298426474, 0.18755362915074925, -0.14308047437526472, -0.2590178575089697, 0.36995137106184356, 0.23830999002106662, -0.06979704075935232, 0.0650928391167193, 0.18519490167516386, 0.3173524550676301, -0.04702390247495576, 0.20419055079559353, 0.22337155858930619, 0.0614614199963621, 0.3123514948576895, 0.18494578758046187, 0.05812721641821871, 0.16537192812044688, 0.31475183638196724, -0.03064451612903212, 0.3023484313487976, 0.07493727972380064, 0.0996705147919488, 0.013259206774573897, 0.3719519890260631, 0.2268096601886579, 0.33103653103653097, 0.28337953278443584, 0.20885350992084475, -0.09031818955249278, -0.11849759142000185, -0.219660479579127, 0.2835580005001023, 0.10742775944096193, 0.048344775232688535, 0.15612557979315703, 0.054847352464217694, -0.3655234411179819, 0.2593523387766398, 0.14821092278719414, 0.0209837473362805, 0.15890585241730293, 0.32145085858125483, 0.13524421649462237, 0.013788916411676138, 0.11773080874798171, 0.2160548143449928, -0.04226869289088544, 0.31211679996808755, 0.18023201827422478, 0.2846885175196416, -0.18037505927178585, 0.26060684985024096})

	fmt.Println("original")
	fmt.Println(sampler.mu)
	fmt.Println(sampler.sigma)
	fmt.Println(Q2(*sampler.data))

	oa := NewOutcomeAggregator(.01/100., 200)

	outcome := 0.0
	stable := false
	for i := 0; i < 100_000_000; i++ {
		outcome = sampler.Sample()
		stable = oa.AddOutcome(outcome)
		if stable {
			fmt.Printf("\nstopped at %d\n", i+1)
			break
		}
	}

	fmt.Println("\nreconstruction (new)")
	fmt.Println(oa.mean)
	fmt.Println(math.Sqrt(oa.variance))
	fmt.Println(oa.q2)
}
