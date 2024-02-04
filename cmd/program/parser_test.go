package program

import (
	"strings"
	"testing"
)

func TestParseEmptyEmail(t *testing.T) {
	p := Parser{}
	fileName := "../../samples/empty.txt"
	em, err := p.Parse(fileName)

	if err != nil {
		t.Fatalf("Error parsing an empty email file! %v", err)
	}
	if em != nil {
		t.Fatalf("got %q, expected nil", em)
	}
}

func TestParaseFullBody(t *testing.T) {
	p := Parser{}
	fileName := "../../samples/email1.txt"
	ex := "\n\nIn today's Daily Update you'll find free reports on\nAmerica Online (AOL), Divine Interventures (DVIN),\nand 3M (MMM); reports on the broadband space, Latin\nAmerican telecom, and more.\n\nFor free research, editor's picks, and more come to the Daily Investor:\nhttp://www.multexinvestor.com/AF004627/magazinecover.asp?promo=unl&d=20001214#\ninvestor\n\n***************************************************************\nYou are receiving this mail because you have registered for\nMultex Investor. To unsubscribe, see bottom of this message.\n***************************************************************\n\n======================== Sponsored by =========================\nWould you own just the energy stocks in the S&P 500?\nSelect Sector SPDRs divides the S&P 500 into nine sector index funds.\nPick and choose just the pieces of the S&P 500 you like best.\nhttp://www.spdrindex.com\n===============================================================\n\nFeatured in today's edition of the Daily Update:\n\n1. SPECIAL ANNOUNCEMENT: Treat yourself to Multex Investor's NEW Personal\nFinance Channel to take advantage of top-notch content and tools  FREE.\n\n2. DAILY FREE SPONSOR REPORT: Robertson Stephens maintains a \"buy\" rating\non Divine Interventures (DVIN).\n\n3. FREE RESEARCH REPORT: Jefferies & Co. rates America Online (AOL) a\n\"buy,\" saying projected growth remains in place.\n\n4. ASK THE ANALYST: Morgan Stanley Dean Witter's Lew Smith in the Analyst\nCorner\n\n5. HOT REPORT: Oscar Gruss & Son's most recent issue of its Broadband\nBrief reports the latest developments in the broadband space.\n\n6. EDITOR'S PICK: Bear Stearns measures the impact of broadband and the\nInternet on telecom in Latin America.\n\n7. FREE STOCK SNAPSHOT: The current analysts' consensus rates 3M (MMM), a\n\"buy/hold.\"\n\n8. JOIN THE MARKETBUZZ: where top financial industry professionals answer\nyour questions and offer insights every market day from noon 'til 2:00\np.m. ET.\n\n9. TRANSCRIPTS FROM WALL STREET: Ash Rajan, senior vice president and\nmarket analyst with Prudential Securities, answers questions about the\nmarket.\n\n======================== Sponsored by =========================\nProfit From AAII's \"Cash Rich\" Stock Screen - 46% YTD Return\n\nWith so much market volatility, how did AAII's \"Cash Rich\"\nStock Screen achieve such stellar returns?  Find the answer by\ntaking a free trial membership from the American Association\nof Individual Investors and using our FREE Stock Screen service at:\nhttp://subs.aaii.com/c/go/XAAI/MTEX1B-aaiitU1?s=S900\n===============================================================\n\n1. NEW ON MULTEX INVESTOR\nTake charge of your personal finances\n\nDo you have endless hours of free time to keep your financial house in\norder? We didn't think so. That's why you need to treat yourself to Multex\nInvestor's NEW Personal Finance Channel to take advantage of top-notch\ncontent and tools  FREE.\nClick here for more information.\nhttp://www.multexpf.com?mktg=sgpftx4&promo=unl&t=10&d=20001214\n\n\n2. DAILY FREE SPONSOR REPORT\nDivine Interventures (DVIN)\n\nRobertson Stephens maintains a \"buy\" rating on Divine Interventures, an\nincubator focused on infrastructure services and business-to-business\n(B2B) exchanges. Register for Robertson Stephens' free-research trial to\naccess this report.\nClick here.\nhttp://www.multexinvestor.com/Download.asp?docid=5018549&sid=9&promo=unl&t=12&\nd=20001214\n\n\n3. FREE RESEARCH REPORT\nHold 'er steady -- America Online (AOL)\n\nAOL's projected growth and proposed merger with Time Warner (TWX) both\nremain in place, says Jefferies & Co., which maintains a \"buy\" rating on\nAOL. In the report, which is free for a limited time, analysts are\nconfident the deal will close soon.\nClick here.\nhttp://www.multexinvestor.com/AF004627/magazinecover.asp?promo=unl&t=11&d=2000\n1214\n\n\n4. TODAY IN THE ANALYST CORNER\nFollowing market trends\n\nMorgan Stanley Dean Witter's Lew Smith sees strong underlying trends\nguiding future market performance. What trends does he point to, and what\nstocks and sectors does he see benefiting from his premise?\n\nHere is your opportunity to gain free access to Morgan Stanley's research.\nSimply register and submit a question below. You will then have a free\ntrial membership to this top Wall Street firms' research!  Lew Smith will\nbe in the Analyst Corner only until 5 p.m. ET Thurs., Dec. 14, so be sure\nto ask your question now.\nAsk the analyst.\nhttp://www.multexinvestor.com/ACHome.asp?promo=unl&t=1&d=20001214\n\n\n5. WHAT'S HOT? RESEARCH REPORTS FROM MULTEX INVESTOR'S HOT LIST\nBreaking the bottleneck -- An update on the broadband space\n\nOscar Gruss & Son's most recent issue of its Broadband Brief reports the\nlatest developments in the broadband space, with coverage of Adaptive\nBroadband (ADAP), Broadcom (BRCM), Efficient Networks (EFNT), and others\n(report for purchase - $25).\nClick here.\nhttp://www.multexinvestor.com/Download.asp?docid=5149041&promo=unl&t=4&d=20001\n214\n\n======================== Sponsored by =========================\nGet Red Herring insight into hot IPOs, investing strategies,\nstocks to watch, future technologies, and more. FREE\nE-newsletters from Redherring.com provide more answers,\nanalysis and opinion to help you make more strategic\ninvesting decisions. Subscribe today\nhttp://www.redherring.com/jump/om/i/multex/email2/subscribe/47.html\n===============================================================\n\n6. EDITOR'S PICK: CURRENT RESEARCH FROM THE CUTTING EDGE\nQue pasa? -- Predicting telecom's future in Latin America\n\nBear Stearns measures the impact of broadband and the Internet on telecom\nin Latin America, saying incumbent local-exchange carriers (ILECs) are\nideally positioned to benefit from the growth of Internet and data\nservices (report for purchase - $150).\nClick here.\nhttp://www.multexinvestor.com/Download.asp?docid=5140995&promo=unl&t=8&d=20001\n214\n\n\n7. FREE STOCK SNAPSHOT\n3M (MMM)\n\nThe current analysts' consensus rates 3M, a \"buy/hold.\" Analysts expect\nthe industrial product manufacturer to earn $4.76 per share in 2000 and\n$5.26 per share in 2001.\nClick here.\nhttp://www.multexinvestor.com/Download.asp?docid=1346414&promo=unl&t=3&d=20001\n214\n\n\n8. JOIN THE MARKETBUZZ!\nCheck out SageOnline\n\nwhere top financial industry professionals answer your questions and offer\ninsights every market day from noon 'til 2:00 p.m. ET.\nClick here.\nhttp://multexinvestor.sageonline.com/page2.asp?id=9512&ps=1&s=2&mktg=evn&promo\n=unl&t=24&d=20001214\n\n\n9. TRANSCRIPTS FROM WALL STREET'S GURUS\nPrudential Securities' Ash Rajan\n\nIn this SageOnline transcript from a chat that took place earlier this\nweek, Ash Rajan, senior vice president and market analyst with Prudential\nSecurities, answers questions about tech, retail, finance, and the outlook\nfor the general market.\nClick here.\nhttp://multexinvestor.sageonline.com/transcript.asp?id=10403&ps=1&s=8&mktg=trn\n&promo=unl&t=13&d=20001214\n\n===================================================================\nPlease send your questions and comments to mailto:investor.help@multex.com\n\nIf you'd like to learn more about Multex Investor, please visit:\nhttp://www.multexinvestor.com/welcome.asp\n\nIf you can't remember your password and/or your user name, click here:\nhttp://www.multexinvestor.com/lostinfo.asp\n\nIf you want to update your email address, please click on the url below:\nhttp://www.multexinvestor.com/edituinfo.asp\n===================================================================\nTo remove yourself from the mailing list for the Daily Update, please\nREPLY to THIS email message with the word UNSUBSCRIBE in the subject\nline. To remove yourself from all Multex Investor mailings, including\nthe Daily Update and The Internet Analyst, please respond with the\nwords NO EMAIL in the subject line.\n\nYou may also unsubscribe on the account update page at:\nhttp://www.multexinvestor.com/edituinfo.asp\n===================================================================\nPlease email advertising inquiries to us at mailto:advertise@multex.com.\n\nFor information on becoming an affiliate click here: \nhttp://www.multexinvestor.com/Affiliates/home.asp?promo=unl\n\nBe sure to check out one of our other newsletters, The Internet Analyst by\nMultex.com. The Internet Analyst informs, educates, and entertains you with\nusable investment data, ideas, experts, and info about the Internet industry.\nTo see this week's issue, click here: http://www.theinternetanalyst.com\n\nIf you are not 100% satisfied with a purchase you make on Multex\nInvestor, we will refund your money."
	em, err := p.Parse(fileName)
	if err != nil {
		t.Fatalf("Error parsing the file, %v", err)
	}

	if strings.Compare(em.Body, ex) != 0 {
		t.Fatalf("got %q, expected %s", em.Body, ex)
	}
}

func Test26(t *testing.T) {
	p := Parser{}
	fileName := "../../samples/26."

	expectedEmail := Email{
		MessageID:               "<15164543.1075855378954.JavaMail.evans@thyme>",
		Date:                    "Wed, 25 Apr 2001 16:52:00 -0700 (PDT)",
		From:                    "phillip.allen@enron.com",
		To:                      []string{"john.lavorato@enron.com"},
		BCC:                     []string{},
		Subject:                 "Re: This morning's Commission meeting delayed",
		MimeVersion:             "1.0",
		ContentType:             "text/plain; charset=us-ascii",
		ContentTransferEncoding: "7bit",
		XFrom:                   "Phillip K Allen",
		XTo:                     []string{"John J Lavorato <John J Lavorato/Corp/Enron@ECT>"},
		Xcc:                     []string{},
		Xbcc:                    []string{},
		XFolder:                 `\Phillip_Allen_Jan2002_1\Allen, Phillip K.\'Sent Mail`,
		XOrigin:                 "Allen-P",
		XFileName:               "pallen (Non-Privileged).pst",
		Body: `


---------------------- Forwarded by Phillip K Allen/HOU/ECT on 04/25/2001 01:51 PM ---------------------------


Ray Alvarez@ENRON
04/25/2001 11:48 AM
To:	Phillip K Allen/HOU/ECT@ECT
cc:	 
Subject:	Re: This morning's Commission meeting delayed   

Phil,  I suspect that discussions/negotiations are taking place behind closed doors "in smoke filled rooms", if not directly between Commissioners then among FERC staffers.  Never say never, but I think it is highly unlikely that the final order will contain a fixed price cap.  I base this belief in large part on what I heard at a luncheon I attended yesterday afternoon at which the keynote speaker was FERC Chairman Curt Hebert.  Although the Chairman began his presentation by expressly stating that he would not comment or answer questions on pending proceedings before the Commission, Hebert had some enlightening comments which relate to price caps:

Price caps are almost never the right answer
Price Caps will have the effect of prolonging shortages
Competitive choices for consumers is the right answer
Any solution, however short term, that does not increase supply or reduce demand, is not acceptable
Eight out of eleven western Governors oppose price caps, in that they would export California's problems to the West

This is the latest intelligence I have on the matter, and it's a pretty strong anti- price cap position.  Of course, Hebert is just one Commissioner out of 3 currently on the Commission, but he controls the meeting agenda and if the draft order is not to his liking, the item could be bumped off the agenda.  Hope this info helps.  Ray




Phillip K Allen@ECT
04/25/2001 02:28 PM
To:	Ray Alvarez/NA/Enron@ENRON
cc:	 

Subject:	Re: This morning's Commission meeting delayed   

Are there behind closed doors discussions being held prior to the meeting?  Is there the potential for a surprise announcement of some sort of fixed price gas or power cap once the open meeting finally happens?






<Embedded StdOleLink>
<Embedded StdOleLink>`,
	}

	em, err := p.Parse(fileName)
	if err != nil {
		t.Fatalf("Error parsing the file, %v", err)
	}

	if strings.Compare(em.MessageID, expectedEmail.MessageID) != 0 {
		t.Fatalf("got %s, expected %s.", em.MessageID, expectedEmail.MessageID)
	}

	if strings.Compare(em.Date, expectedEmail.Date) != 0 {
		t.Fatalf("got %s, expected %s.", em.Date, expectedEmail.Date)
	}

	if strings.Compare(em.From, expectedEmail.From) != 0 {
		t.Fatalf("got %s, expected %s.", em.From, expectedEmail.From)
	}

	if !testEq(em.To, expectedEmail.To) {
		t.Fatalf("got %v, expected %v.", em.To, expectedEmail.To)
	}

	if strings.Compare(em.Body, expectedEmail.Body) != 0 {
		t.Fatalf("got %s, expected %s.", em.Body, expectedEmail.Body)
	}

}

func testEq(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestSubject(t *testing.T) {
	p := Parser{}
	fileName := "../../samples/17.txt"
	expectedSubject := `Hewlett Packard Conference call on Wireless and Handheld
 Technologies, December 14th, 1:30-2:30 PM`
	em, err := p.Parse(fileName)
	if err != nil {
		t.Fatalf("Error parsing the file, %v", err)
	}

	if strings.Compare(em.Subject, expectedSubject) != 0 {
		t.Fatalf("got %q, expected %q", em.Subject, expectedSubject)
	}
}
