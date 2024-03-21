package main

/*
@Author Aidan Scott
Stores article data
*/
var localArticles = map[string][2]string{
	"updates": {"Belt sector AP-089 daily updates",
		"From: Jimle Brown, news officer of AP-089, 10/10/2224 SST\n" +
			"Hello everyone out there in AP-089! hope your solar days have been full of profitable mining ventures!\n" +
			"I heard Tom had his most profitable solar year, with the recent uptick of mega projects under construction \n" +
			"across the settled systems it makes sense that our products would become more valuable. So everyone, you too Sarah, \n" +
			"can look forward to increased prices at the carrier fleets. Now for our monthly updates price jumps to watch out for.\n" +
			"-> Iron 10.78/T > 11.05\n" +
			"-> Water 104.77/T > 160.01\n" +
			"-> Gold 30.89/T > 30.10\n" +
			"-> Generals 2.05 > 2.06\n" +
			"Prices rise, and not just mining products, we have a lot of new neighbors outside our little sector. \n" +
			"Chi from AP-090 told me that 5 new miners joined the sector just over this past solar month! We could see \n" +
			"an increase to our population from 10 to ... maybe 11! \n" +
			"That's all for today's updates! Stay calm and keep mining!\n"},
	"comets": {"Mining asteroids? Try comets instead!",
		"From: Sarah's Radio show, transcript summary version, 10/9/2224 SST\n" +
			"Hey there all you 10 or so out there listening to me on the void today. I want to talk about an alternative mining source. \n" +
			"We mine rocks, dusty rocks and dust all day every day, and that has done us well, but there is another resource \n" +
			"that we have seen an increase in price, to most of our dismays... WATER! By ton water has increased from 104.77 to 160.01 YenDollars \n" +
			"in the past solar month! That is a massive increase, but we all know that especially when the carriers come by they try to sell us \n" +
			"it at an even higher price! Well, I have a solution for you, our sector has a varying amount of comets or icy asteroids, around 400-500 \n" +
			"at any given time. I started mining these resources and I have never been happier! I am swindling the carriers now! I just sold a couple tons of \n" +
			"H2O for 170.01 YenDollars. \n" +
			"I hope to see you farming that ice, see you next time!\n"},
	"carriers": {"WARNING: local carriers will be going on break next solar week",
		"From: Carrier's Guild notification agency, 10/8/2224 SST\n" +
			"#RELAYED ON LOCAL NETWORKS ACROSS THE BELT ON 10/9/2224 SOLAR STANDARD TIME#" +
			"---WARNING--- carriers will be taking a vacation this next week for Victory Week, celebrating the victory we share during \n" +
			"the Great Solar War. We hope that you take a break as well and appreciate what we have built together.\n" +
			"After Victory Week the guild will have a couple more ships in the AP-08* systems providing more opportunity to sell your product.\n" +
			"Remember to greet your local Carriers with a smile! They are only there to help you prosper!"},
}

// returns map of articles
func getLocalNews() map[string][2]string {
	return localArticles
}
