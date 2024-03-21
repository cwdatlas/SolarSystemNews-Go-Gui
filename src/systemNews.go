package main

/*
@Author Aidan Scott
Stores article data
*/
var systemArticles = map[string][2]string{
	"venus": {"New Life on Venus",
		"From: Venus immigration agency, 10/8/2224 SST\n" +
			"Come to lovely Venus with is great sky cities like Vaccos, and Tenebra!\n" +
			"50 kilometers over the surface of Venus is the best place to live in the System!\n" +
			"No better place to raise a family than the sunny, warm days in the Vaccos suburbs!\n" +
			"Come today and reserve a high in demand home! People come every day, space is running out!\n" +
			"Come today!",
	},
	"mirror": {"New Mirror Project in Ganymede orbit",
		"From: Juvian Progress and Peoples Council, 10/3/2224 SST.\n" +
			"A new mega project is being built in orbit of IO. This project is meant to alter the local day length.\n" +
			"The mirror is also meant to warm the surface to a more barable tempurature in preparation for larger \n" +
			"terraforming projects. Our largest moon will soon be provided with adequate power to run large refineries \n" +
			"making Ganymede the core of the settled system's manufacturing capabilities. This is only our first step \n" +
			"towards greatness for the Juvian system, more projects on the way. Гарного вам дня"},
	"work": {"Work on the Moon; What you need to know",
		"From The Luna Post, Shackleton, 10/2/2224 SST\n" +
			"The Moon is the fastest growing sector in the settled systems, good work being the primary reason for that.\n" +
			"The Space Docks all throughout the Terran system consume the vast majority of our mined and manufactured materials.\n" +
			"To power The Terran military we must mine, and process more! The Democratic United System is providing pools of \n" +
			"YenDollars to anyone who wants to move to The Moon and start your own company or work for the big three, \n" +
			"Southern NewPeru Mining, Saudi Refco, and Yamato Tritium. The big three are great employers with high wages and \n" +
			"safe work conditions compared to IO and Belt mining and refining. If your more into manufacturing then Hanwha Space \n" +
			"provides great vessel construction in Lunar orbit. They produce 40% of all the ships in the Terran system and 20% in all \n" +
			"the settled systems. If you want to help our great military, then Lockheed X Martin could be your best bet. You can commute \n" +
			"from The Moon to GEO every month or so. Its a little distant, but it makes the best money on luna except for Yamato Tritium atomic \n" +
			"engineers. Overall, there are many opportunities on The Moon in the Terran System so join us pushing the DUS forward!"}}

// returns map of articles
func getSystemNews() map[string][2]string {
	return systemArticles
}
