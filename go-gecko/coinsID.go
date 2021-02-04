package main

import (
	"encoding/json"
	"fmt"
	"github.com/pubgo/xerror"
	"log"

	gecko "github.com/superoo7/go-gecko/v3"
)

func main() {
	cg := gecko.NewClient(nil)
	coin, err := cg.CoinsID("dogecoin", true, true, true, true, true, true)
	if err != nil {
		log.Fatal(err)
	}

	dt, err := json.MarshalIndent(coin, "", "\t")
	xerror.Panic(err)
	fmt.Println(string(dt))
}

var _ = `{
        "id": "dogecoin",
        "symbol": "doge",
        "name": "Dogecoin",
        "block_time_in_minutes": 1,
        "categories": [
                "Cryptocurrency"
        ],
        "localization": {
                "ar": "الدوجكوين",
                "de": "Dogecoin",
                "en": "Dogecoin",
                "es": "Dogecoin",
                "fr": "Dogecoin",
                "hu": "Dogecoin",
                "id": "Dogecoin",
                "it": "Dogecoin",
                "ja": "ドージコイン",
                "ko": "도지코인",
                "nl": "Dogecoin",
                "pl": "Dogecoin",
                "pt": "Dogecoin",
                "ro": "Dogecoin",
                "ru": "Dogecoin",
                "sv": "Dogecoin",
                "th": "Dogecoin",
                "tr": "Dogecoin",
                "vi": "Dogecoin",
                "zh": "狗狗币",
                "zh-tw": "狗狗幣"
        },
        "description": {
                "ar": "",
                "de": "Dogecoin is a cryptocurrency based on the popular \"Doge\" Internet meme and features a Shiba Inu on its logo. Dogecoin is a \u003ca href=\"/coins/litecoin\"\u003eLitecoin\u003c/a\u003e fork. Introduced as a \"joke currency\" on 6 December 2013, Dogecoin quickly developed its own online community and reached a capitalization of US$60 million in January 2014. Compared with other cryptocurrencies, Dogecoin had a fast initial coin production schedule: 100 billion coins were in circulation by mid-2015, with an additional 5.256 billion coins every year thereafter. As of 30 June 2015, the 100 billionth Dogecoin had been mined. \r\n\r\nDogecoin was created by Billy Markus from Portland, Oregon and  Jackson Palmer from Sydney, Australia. Both wanted to create a fun cryptocurrency that will appeal beyond the core Bitcoin audience. Dogecoin is primarily used as a tipping system on Reddit and Twitter where users tip each other for creating or sharing good content. The community is very active in organising fundraising activities for deserving causes.\r\n\r\nThe developers of Dogecoin haven’t made any major changes to the coin since 2015. This means that Dogecoin could get left behind and is why Shibas are leaving Dogecoin to join more advanced platforms like Ethereum. One of Dogecoin strengths is its relaxed and fun-loving community. However, this is also a weakness because other currencies are way more professional.\r\n\r\nTo purchase Dogecoin, it involves downloading a crypto wallet, setting up a crypto exchange account and then trading away for your desired crypto currency. Once we have set up an account with a DOGE currency exchange and deposited some funds, you are ready to start trading. ",
                "en": "Dogecoin is a cryptocurrency based on the popular \"Doge\" Internet meme and features a Shiba Inu on its logo. Dogecoin is a \u003ca href=\"/coins/litecoin\"\u003eLitecoin\u003c/a\u003e fork. Introduced as a \"joke currency\" on 6 December 2013, Dogecoin quickly developed its own online community and reached a capitalization of US$60 million in January 2014. Compared with other cryptocurrencies, Dogecoin had a fast initial coin production schedule: 100 billion coins were in circulation by mid-2015, with an additional 5.256 billion coins every year thereafter. As of 30 June 2015, the 100 billionth Dogecoin had been mined. \r\n\r\nDogecoin was created by Billy Markus from Portland, Oregon and  Jackson Palmer from Sydney, Australia. Both wanted to create a fun cryptocurrency that will appeal beyond the core Bitcoin audience. Dogecoin is primarily used as a tipping system on Reddit and Twitter where users tip each other for creating or sharing good content. The community is very active in organising fundraising activities for deserving causes.\r\n\r\nThe developers of Dogecoin haven’t made any major changes to the coin since 2015. This means that Dogecoin could get left behind and is why Shibas are leaving Dogecoin to join more advanced platforms like Ethereum. One of Dogecoin strengths is its relaxed and fun-loving community. However, this is also a weakness because other currencies are way more professional.\r\n\r\nTo purchase Dogecoin, it involves downloading a crypto wallet, setting up a crypto exchange account and then trading away for your desired crypto currency. Once we have set up an account with a DOGE currency exchange and deposited some funds, you are ready to start trading. ",
                "es": "Dogecoin is a cryptocurrency based on the popular \"Doge\" Internet meme and features a Shiba Inu on its logo. Dogecoin is a \u003ca href=\"/coins/litecoin\"\u003eLitecoin\u003c/a\u003e fork. Introduced as a \"joke currency\" on 6 December 2013, Dogecoin quickly developed its own online community and reached a capitalization of US$60 million in January 2014. Compared with other cryptocurrencies, Dogecoin had a fast initial coin production schedule: 100 billion coins were in circulation by mid-2015, with an additional 5.256 billion coins every year thereafter. As of 30 June 2015, the 100 billionth Dogecoin had been mined. \r\n\r\nDogecoin was created by Billy Markus from Portland, Oregon and  Jackson Palmer from Sydney, Australia. Both wanted to create a fun cryptocurrency that will appeal beyond the core Bitcoin audience. Dogecoin is primarily used as a tipping system on Reddit and Twitter where users tip each other for creating or sharing good content. The community is very active in organising fundraising activities for deserving causes.\r\n\r\nThe developers of Dogecoin haven’t made any major changes to the coin since 2015. This means that Dogecoin could get left behind and is why Shibas are leaving Dogecoin to join more advanced platforms like Ethereum. One of Dogecoin strengths is its relaxed and fun-loving community. However, this is also a weakness because other currencies are way more professional.\r\n\r\nTo purchase Dogecoin, it involves downloading a crypto wallet, setting up a crypto exchange account and then trading away for your desired crypto currency. Once we have set up an account with a DOGE currency exchange and deposited some funds, you are ready to start trading. ",
                "fr": "Dogecoin is a cryptocurrency based on the popular \"Doge\" Internet meme and features a Shiba Inu on its logo. Dogecoin is a \u003ca href=\"/coins/litecoin\"\u003eLitecoin\u003c/a\u003e fork. Introduced as a \"joke currency\" on 6 December 2013, Dogecoin quickly developed its own online community and reached a capitalization of US$60 million in January 2014. Compared with other cryptocurrencies, Dogecoin had a fast initial coin production schedule: 100 billion coins were in circulation by mid-2015, with an additional 5.256 billion coins every year thereafter. As of 30 June 2015, the 100 billionth Dogecoin had been mined. \r\n\r\nDogecoin was created by Billy Markus from Portland, Oregon and  Jackson Palmer from Sydney, Australia. Both wanted to create a fun cryptocurrency that will appeal beyond the core Bitcoin audience. Dogecoin is primarily used as a tipping system on Reddit and Twitter where users tip each other for creating or sharing good content. The community is very active in organising fundraising activities for deserving causes.\r\n\r\nThe developers of Dogecoin haven’t made any major changes to the coin since 2015. This means that Dogecoin could get left behind and is why Shibas are leaving Dogecoin to join more advanced platforms like Ethereum. One of Dogecoin strengths is its relaxed and fun-loving community. However, this is also a weakness because other currencies are way more professional.\r\n\r\nTo purchase Dogecoin, it involves downloading a crypto wallet, setting up a crypto exchange account and then trading away for your desired crypto currency. Once we have set up an account with a DOGE currency exchange and deposited some funds, you are ready to start trading. ",
                "hu": "Dogecoin is a cryptocurrency based on the popular \"Doge\" Internet meme and features a Shiba Inu on its logo. Dogecoin is a \u003ca href=\"/coins/litecoin\"\u003eLitecoin\u003c/a\u003e fork. Introduced as a \"joke currency\" on 6 December 2013, Dogecoin quickly developed its own online community and reached a capitalization of US$60 million in January 2014. Compared with other cryptocurrencies, Dogecoin had a fast initial coin production schedule: 100 billion coins were in circulation by mid-2015, with an additional 5.256 billion coins every year thereafter. As of 30 June 2015, the 100 billionth Dogecoin had been mined. \r\n\r\nDogecoin was created by Billy Markus from Portland, Oregon and  Jackson Palmer from Sydney, Australia. Both wanted to create a fun cryptocurrency that will appeal beyond the core Bitcoin audience. Dogecoin is primarily used as a tipping system on Reddit and Twitter where users tip each other for creating or sharing good content. The community is very active in organising fundraising activities for deserving causes.\r\n\r\nThe developers of Dogecoin haven’t made any major changes to the coin since 2015. This means that Dogecoin could get left behind and is why Shibas are leaving Dogecoin to join more advanced platforms like Ethereum. One of Dogecoin strengths is its relaxed and fun-loving community. However, this is also a weakness because other currencies are way more professional.\r\n\r\nTo purchase Dogecoin, it involves downloading a crypto wallet, setting up a crypto exchange account and then trading away for your desired crypto currency. Once we have set up an account with a DOGE currency exchange and deposited some funds, you are ready to start trading. ",
                "id": "Dogecoin is a cryptocurrency based on the popular \"Doge\" Internet meme and features a Shiba Inu on its logo. Dogecoin is a \u003ca href=\"/coins/litecoin\"\u003eLitecoin\u003c/a\u003e fork. Introduced as a \"joke currency\" on 6 December 2013, Dogecoin quickly developed its own online community and reached a capitalization of US$60 million in January 2014. Compared with other cryptocurrencies, Dogecoin had a fast initial coin production schedule: 100 billion coins were in circulation by mid-2015, with an additional 5.256 billion coins every year thereafter. As of 30 June 2015, the 100 billionth Dogecoin had been mined. \r\n\r\nDogecoin was created by Billy Markus from Portland, Oregon and  Jackson Palmer from Sydney, Australia. Both wanted to create a fun cryptocurrency that will appeal beyond the core Bitcoin audience. Dogecoin is primarily used as a tipping system on Reddit and Twitter where users tip each other for creating or sharing good content. The community is very active in organising fundraising activities for deserving causes.\r\n\r\nThe developers of Dogecoin haven’t made any major changes to the coin since 2015. This means that Dogecoin could get left behind and is why Shibas are leaving Dogecoin to join more advanced platforms like Ethereum. One of Dogecoin strengths is its relaxed and fun-loving community. However, this is also a weakness because other currencies are way more professional.\r\n\r\nTo purchase Dogecoin, it involves downloading a crypto wallet, setting up a crypto exchange account and then trading away for your desired crypto currency. Once we have set up an account with a DOGE currency exchange and deposited some funds, you are ready to start trading. ",
                "it": "Dogecoin is a cryptocurrency based on the popular \"Doge\" Internet meme and features a Shiba Inu on its logo. Dogecoin is a \u003ca href=\"/coins/litecoin\"\u003eLitecoin\u003c/a\u003e fork. Introduced as a \"joke currency\" on 6 December 2013, Dogecoin quickly developed its own online community and reached a capitalization of US$60 million in January 2014. Compared with other cryptocurrencies, Dogecoin had a fast initial coin production schedule: 100 billion coins were in circulation by mid-2015, with an additional 5.256 billion coins every year thereafter. As of 30 June 2015, the 100 billionth Dogecoin had been mined. \r\n\r\nDogecoin was created by Billy Markus from Portland, Oregon and  Jackson Palmer from Sydney, Australia. Both wanted to create a fun cryptocurrency that will appeal beyond the core Bitcoin audience. Dogecoin is primarily used as a tipping system on Reddit and Twitter where users tip each other for creating or sharing good content. The community is very active in organising fundraising activities for deserving causes.\r\n\r\nThe developers of Dogecoin haven’t made any major changes to the coin since 2015. This means that Dogecoin could get left behind and is why Shibas are leaving Dogecoin to join more advanced platforms like Ethereum. One of Dogecoin strengths is its relaxed and fun-loving community. However, this is also a weakness because other currencies are way more professional.\r\n\r\nTo purchase Dogecoin, it involves downloading a crypto wallet, setting up a crypto exchange account and then trading away for your desired crypto currency. Once we have set up an account with a DOGE currency exchange and deposited some funds, you are ready to start trading. ",
                "ja": "",
                "ko": "처음에 \"joke currency\"라고 불리기도 하면서 장난처럼 시작한 도지코인은 일본 개인 시바 이누를 마스코트로 사용합니다. 이 시바견은 인터넷에서 재미로근한 가상화폐라는 점을 강조합니다. 실제로 당장 홈페이지만 접속해도, 아주 쉽게 도지코인 지갑을 설치할 수 있습니다. \r\n\r\n개발자 빌리 마커스는 불법 마약 거래 사이트 실크r\n\r\n코인 특징\r\n1. 도지코인은 라이트코인에서 포크된 럭키코인에서 포크되었습니다. 그래서 처음에는 럭키코인처럼 채굴보상이 랜덤이었는데, 이후 정해진 보상으로 정책을 바돼있었는데, 무제한 생산으로 바뀌었습니다. 현재 10,000개의 코인이 1분마다 생겨나는 중이고, 1년에는 52억 개의 새로운 도지코인이 생겨납니다. 2015년 6월 30일 1,000억 개의 코얻었습니다. 즉, 도지코인을 이용해 사용자들이 흥미롭거나 가치 있는 콘텐츠를 제공한 사람에게 팁을 주는 것입니다. 레딧, 트위터, 트위치티비(Twitch.TV)등에서 이런 서비스를 제공하는 도지팁봇(Dogetipbot)이 등장하기도 했으나 현재 사용 가능한 팁봇은 제한적입니다. ",
                "nl": "Dogecoin is a cryptocurrency based on the popular \"Doge\" Internet meme and features a Shiba Inu on its logo. Dogecoin is a \u003ca href=\"/coins/litecoin\"\u003eLitecoin\u003c/a\u003e fork. Introduced as a \"joke currency\" on 6 December 2013, Dogecoin quickly developed its own online community and reached a capitalization of US$60 million in January 2014. Compared with other cryptocurrencies, Dogecoin had a fast initial coin production schedule: 100 billion coins were in circulation by mid-2015, with an additional 5.256 billion coins every year thereafter. As of 30 June 2015, the 100 billionth Dogecoin had been mined. \r\n\r\nDogecoin was created by Billy Markus from Portland, Oregon and  Jackson Palmer from Sydney, Australia. Both wanted to create a fun cryptocurrency that will appeal beyond the core Bitcoin audience. Dogecoin is primarily used as a tipping system on Reddit and Twitter where users tip each other for creating or sharing good content. The community is very active in organising fundraising activities for deserving causes.\r\n\r\nThe developers of Dogecoin haven’t made any major changes to the coin since 2015. This means that Dogecoin could get left behind and is why Shibas are leaving Dogecoin to join more advanced platforms like Ethereum. One of Dogecoin strengths is its relaxed and fun-loving community. However, this is also a weakness because other currencies are way more professional.\r\n\r\nTo purchase Dogecoin, it involves downloading a crypto wallet, setting up a crypto exchange account and then trading away for your desired crypto currency. Once we have set up an account with a DOGE currency exchange and deposited some funds, you are ready to start trading. ",
                "pl": "Dogecoin is a cryptocurrency based on the popular \"Doge\" Internet meme and features a Shiba Inu on its logo. Dogecoin is a \u003ca href=\"/coins/litecoin\"\u003eLitecoin\u003c/a\u003e fork. Introduced as a \"joke currency\" on 6 December 2013, Dogecoin quickly developed its own online community and reached a capitalization of US$60 million in January 2014. Compared with other cryptocurrencies, Dogecoin had a fast initial coin production schedule: 100 billion coins were in circulation by mid-2015, with an additional 5.256 billion coins every year thereafter. As of 30 June 2015, the 100 billionth Dogecoin had been mined. \r\n\r\nDogecoin was created by Billy Markus from Portland, Oregon and  Jackson Palmer from Sydney, Australia. Both wanted to create a fun cryptocurrency that will appeal beyond the core Bitcoin audience. Dogecoin is primarily used as a tipping system on Reddit and Twitter where users tip each other for creating or sharing good content. The community is very active in organising fundraising activities for deserving causes.\r\n\r\nThe developers of Dogecoin haven’t made any major changes to the coin since 2015. This means that Dogecoin could get left behind and is why Shibas are leaving Dogecoin to join more advanced platforms like Ethereum. One of Dogecoin strengths is its relaxed and fun-loving community. However, this is also a weakness because other currencies are way more professional.\r\n\r\nTo purchase Dogecoin, it involves downloading a crypto wallet, setting up a crypto exchange account and then trading away for your desired crypto currency. Once we have set up an account with a DOGE currency exchange and deposited some funds, you are ready to start trading. ",
                "pt": "Dogecoin is a cryptocurrency based on the popular \"Doge\" Internet meme and features a Shiba Inu on its logo. Dogecoin is a \u003ca href=\"/coins/litecoin\"\u003eLitecoin\u003c/a\u003e fork. Introduced as a \"joke currency\" on 6 December 2013, Dogecoin quickly developed its own online community and reached a capitalization of US$60 million in January 2014. Compared with other cryptocurrencies, Dogecoin had a fast initial coin production schedule: 100 billion coins were in circulation by mid-2015, with an additional 5.256 billion coins every year thereafter. As of 30 June 2015, the 100 billionth Dogecoin had been mined. \r\n\r\nDogecoin was created by Billy Markus from Portland, Oregon and  Jackson Palmer from Sydney, Australia. Both wanted to create a fun cryptocurrency that will appeal beyond the core Bitcoin audience. Dogecoin is primarily used as a tipping system on Reddit and Twitter where users tip each other for creating or sharing good content. The community is very active in organising fundraising activities for deserving causes.\r\n\r\nThe developers of Dogecoin haven’t made any major changes to the coin since 2015. This means that Dogecoin could get left behind and is why Shibas are leaving Dogecoin to join more advanced platforms like Ethereum. One of Dogecoin strengths is its relaxed and fun-loving community. However, this is also a weakness because other currencies are way more professional.\r\n\r\nTo purchase Dogecoin, it involves downloading a crypto wallet, setting up a crypto exchange account and then trading away for your desired crypto currency. Once we have set up an account with a DOGE currency exchange and deposited some funds, you are ready to start trading. ",
                "ro": "Dogecoin is a cryptocurrency based on the popular \"Doge\" Internet meme and features a Shiba Inu on its logo. Dogecoin is a \u003ca href=\"/coins/litecoin\"\u003eLitecoin\u003c/a\u003e fork. Introduced as a \"joke currency\" on 6 December 2013, Dogecoin quickly developed its own online community and reached a capitalization of US$60 million in January 2014. Compared with other cryptocurrencies, Dogecoin had a fast initial coin production schedule: 100 billion coins were in circulation by mid-2015, with an additional 5.256 billion coins every year thereafter. As of 30 June 2015, the 100 billionth Dogecoin had been mined. \r\n\r\nDogecoin was created by Billy Markus from Portland, Oregon and  Jackson Palmer from Sydney, Australia. Both wanted to create a fun cryptocurrency that will appeal beyond the core Bitcoin audience. Dogecoin is primarily used as a tipping system on Reddit and Twitter where users tip each other for creating or sharing good content. The community is very active in organising fundraising activities for deserving causes.\r\n\r\nThe developers of Dogecoin haven’t made any major changes to the coin since 2015. This means that Dogecoin could get left behind and is why Shibas are leaving Dogecoin to join more advanced platforms like Ethereum. One of Dogecoin strengths is its relaxed and fun-loving community. However, this is also a weakness because other currencies are way more professional.\r\n\r\nTo purchase Dogecoin, it involves downloading a crypto wallet, setting up a crypto exchange account and then trading away for your desired crypto currency. Once we have set up an account with a DOGE currency exchange and deposited some funds, you are ready to start trading. ",
                "ru": "Dogecoin is a cryptocurrency based on the popular \"Doge\" Internet meme and features a Shiba Inu on its logo. Dogecoin is a \u003ca href=\"/coins/litecoin\"\u003eLitecoin\u003c/a\u003e fork. Introduced as a \"joke currency\" on 6 December 2013, Dogecoin quickly developed its own online community and reached a capitalization of US$60 million in January 2014. Compared with other cryptocurrencies, Dogecoin had a fast initial coin production schedule: 100 billion coins were in circulation by mid-2015, with an additional 5.256 billion coins every year thereafter. As of 30 June 2015, the 100 billionth Dogecoin had been mined. \r\n\r\nDogecoin was created by Billy Markus from Portland, Oregon and  Jackson Palmer from Sydney, Australia. Both wanted to create a fun cryptocurrency that will appeal beyond the core Bitcoin audience. Dogecoin is primarily used as a tipping system on Reddit and Twitter where users tip each other for creating or sharing good content. The community is very active in organising fundraising activities for deserving causes.\r\n\r\nThe developers of Dogecoin haven’t made any major changes to the coin since 2015. This means that Dogecoin could get left behind and is why Shibas are leaving Dogecoin to join more advanced platforms like Ethereum. One of Dogecoin strengths is its relaxed and fun-loving community. However, this is also a weakness because other currencies are way more professional.\r\n\r\nTo purchase Dogecoin, it involves downloading a crypto wallet, setting up a crypto exchange account and then trading away for your desired crypto currency. Once we have set up an account with a DOGE currency exchange and deposited some funds, you are ready to start trading. ",
                "sv": "Dogecoin is a cryptocurrency based on the popular \"Doge\" Internet meme and features a Shiba Inu on its logo. Dogecoin is a \u003ca href=\"/coins/litecoin\"\u003eLitecoin\u003c/a\u003e fork. Introduced as a \"joke currency\" on 6 December 2013, Dogecoin quickly developed its own online community and reached a capitalization of US$60 million in January 2014. Compared with other cryptocurrencies, Dogecoin had a fast initial coin production schedule: 100 billion coins were in circulation by mid-2015, with an additional 5.256 billion coins every year thereafter. As of 30 June 2015, the 100 billionth Dogecoin had been mined. \r\n\r\nDogecoin was created by Billy Markus from Portland, Oregon and  Jackson Palmer from Sydney, Australia. Both wanted to create a fun cryptocurrency that will appeal beyond the core Bitcoin audience. Dogecoin is primarily used as a tipping system on Reddit and Twitter where users tip each other for creating or sharing good content. The community is very active in organising fundraising activities for deserving causes.\r\n\r\nThe developers of Dogecoin haven’t made any major changes to the coin since 2015. This means that Dogecoin could get left behind and is why Shibas are leaving Dogecoin to join more advanced platforms like Ethereum. One of Dogecoin strengths is its relaxed and fun-loving community. However, this is also a weakness because other currencies are way more professional.\r\n\r\nTo purchase Dogecoin, it involves downloading a crypto wallet, setting up a crypto exchange account and then trading away for your desired crypto currency. Once we have set up an account with a DOGE currency exchange and deposited some funds, you are ready to start trading. ",
                "th": "Dogecoin is a cryptocurrency based on the popular \"Doge\" Internet meme and features a Shiba Inu on its logo. Dogecoin is a \u003ca href=\"/coins/litecoin\"\u003eLitecoin\u003c/a\u003e fork. Introduced as a \"joke currency\" on 6 December 2013, Dogecoin quickly developed its own online community and reached a capitalization of US$60 million in January 2014. Compared with other cryptocurrencies, Dogecoin had a fast initial coin production schedule: 100 billion coins were in circulation by mid-2015, with an additional 5.256 billion coins every year thereafter. As of 30 June 2015, the 100 billionth Dogecoin had been mined. \r\n\r\nDogecoin was created by Billy Markus from Portland, Oregon and  Jackson Palmer from Sydney, Australia. Both wanted to create a fun cryptocurrency that will appeal beyond the core Bitcoin audience. Dogecoin is primarily used as a tipping system on Reddit and Twitter where users tip each other for creating or sharing good content. The community is very active in organising fundraising activities for deserving causes.\r\n\r\nThe developers of Dogecoin haven’t made any major changes to the coin since 2015. This means that Dogecoin could get left behind and is why Shibas are leaving Dogecoin to join more advanced platforms like Ethereum. One of Dogecoin strengths is its relaxed and fun-loving community. However, this is also a weakness because other currencies are way more professional.\r\n\r\nTo purchase Dogecoin, it involves downloading a crypto wallet, setting up a crypto exchange account and then trading away for your desired crypto currency. Once we have set up an account with a DOGE currency exchange and deposited some funds, you are ready to start trading. ",
                "tr": "Dogecoin is a cryptocurrency based on the popular \"Doge\" Internet meme and features a Shiba Inu on its logo. Dogecoin is a \u003ca href=\"/coins/litecoin\"\u003eLitecoin\u003c/a\u003e fork. Introduced as a \"joke currency\" on 6 December 2013, Dogecoin quickly developed its own online community and reached a capitalization of US$60 million in January 2014. Compared with other cryptocurrencies, Dogecoin had a fast initial coin production schedule: 100 billion coins were in circulation by mid-2015, with an additional 5.256 billion coins every year thereafter. As of 30 June 2015, the 100 billionth Dogecoin had been mined. \r\n\r\nDogecoin was created by Billy Markus from Portland, Oregon and  Jackson Palmer from Sydney, Australia. Both wanted to create a fun cryptocurrency that will appeal beyond the core Bitcoin audience. Dogecoin is primarily used as a tipping system on Reddit and Twitter where users tip each other for creating or sharing good content. The community is very active in organising fundraising activities for deserving causes.\r\n\r\nThe developers of Dogecoin haven’t made any major changes to the coin since 2015. This means that Dogecoin could get left behind and is why Shibas are leaving Dogecoin to join more advanced platforms like Ethereum. One of Dogecoin strengths is its relaxed and fun-loving community. However, this is also a weakness because other currencies are way more professional.\r\n\r\nTo purchase Dogecoin, it involves downloading a crypto wallet, setting up a crypto exchange account and then trading away for your desired crypto currency. Once we have set up an account with a DOGE currency exchange and deposited some funds, you are ready to start trading. ",
                "vi": "Dogecoin is a cryptocurrency based on the popular \"Doge\" Internet meme and features a Shiba Inu on its logo. Dogecoin is a \u003ca href=\"/coins/litecoin\"\u003eLitecoin\u003c/a\u003e fork. Introduced as a \"joke currency\" on 6 December 2013, Dogecoin quickly developed its own online community and reached a capitalization of US$60 million in January 2014. Compared with other cryptocurrencies, Dogecoin had a fast initial coin production schedule: 100 billion coins were in circulation by mid-2015, with an additional 5.256 billion coins every year thereafter. As of 30 June 2015, the 100 billionth Dogecoin had been mined. \r\n\r\nDogecoin was created by Billy Markus from Portland, Oregon and  Jackson Palmer from Sydney, Australia. Both wanted to create a fun cryptocurrency that will appeal beyond the core Bitcoin audience. Dogecoin is primarily used as a tipping system on Reddit and Twitter where users tip each other for creating or sharing good content. The community is very active in organising fundraising activities for deserving causes.\r\n\r\nThe developers of Dogecoin haven’t made any major changes to the coin since 2015. This means that Dogecoin could get left behind and is why Shibas are leaving Dogecoin to join more advanced platforms like Ethereum. One of Dogecoin strengths is its relaxed and fun-loving community. However, this is also a weakness because other currencies are way more professional.\r\n\r\nTo purchase Dogecoin, it involves downloading a crypto wallet, setting up a crypto exchange account and then trading away for your desired crypto currency. Once we have set up an account with a DOGE currency exchange and deposited some funds, you are ready to start trading. ",
                "zh": "",
                "zh-tw": ""
        },
        "links": {
                "announcement_url": [
                        "",
                        ""
                ],
                "bitcointalk_thread_identifier": 361813,
                "blockchain_site": [
                        "https://blockchair.com/dogecoin",
                        "https://doge.tokenview.com/",
                        "http://dogechain.info/chain/Dogecoin",
                        "",
                        ""
                ],
                "chat_url": [
                        "",
                        "",
                        ""
                ],
                "facebook_username": "OfficialDogecoin",
                "homepage": [
                        "http://dogecoin.com/",
                        "",
                        ""
                ],
                "official_forum_url": [
                        "http://discuss.dogecoin.com/",
                        "",
                        ""
                ],
                "repos_url": {
                        "bitbucket": [],
                        "github": [
                                "https://github.com/dogecoin/dogecoin",
                                "https://github.com/dogecoin/libdohj"
                        ]
                },
                "subreddit_url": "https://www.reddit.com/r/dogecoin/",
                "telegram_channel_identifier": "",
                "twitter_screen_name": "dogecoin"
        },
        "image": {
                "thumb": "https://assets.coingecko.com/coins/images/5/thumb/dogecoin.png?1547792256",
                "small": "https://assets.coingecko.com/coins/images/5/small/dogecoin.png?1547792256",
                "large": "https://assets.coingecko.com/coins/images/5/large/dogecoin.png?1547792256"
        },
        "country_origin": "",
        "genesis_date": "2013-12-08",
        "market_cap_rank": 14,
        "coingecko_rank": 4,
        "coingecko_score": 72.642,
        "developer_score": 75.512,
        "community_score": 90.109,
        "liquidity_score": 86.965,
        "public_interest_score": 0,
        "market_data": {
                "current_price": {
                        "aed": 0.162595,
                        "ars": 3.89,
                        "aud": 0.058282,
                        "bch": 0.00010554,
                        "bdt": 3.75,
                        "bhd": 0.01668876,
                        "bits": 1.21,
                        "bmd": 0.04426774,
                        "bnb": 0.00082535,
                        "brl": 0.239849,
                        "btc": 0.00000121,
                        "cad": 0.056804,
                        "chf": 0.03998404,
                        "clp": 32.74,
                        "cny": 0.286492,
                        "czk": 0.956977,
                        "dkk": 0.274857,
                        "dot": 0.00223111,
                        "eos": 0.01450187,
                        "eth": 0.00002782,
                        "eur": 0.03696338,
                        "gbp": 0.03239925,
                        "hkd": 0.343185,
                        "huf": 13.17,
                        "idr": 622.23,
                        "ils": 0.145703,
                        "inr": 3.23,
                        "jpy": 4.67,
                        "krw": 49.6,
                        "kwd": 0.01341268,
                        "link": 0.00187284,
                        "lkr": 8.56,
                        "ltc": 0.00030546,
                        "mmk": 61.41,
                        "mxn": 0.906166,
                        "myr": 0.179661,
                        "ngn": 16.82,
                        "nok": 0.38245,
                        "nzd": 0.061827,
                        "php": 2.13,
                        "pkr": 7.09,
                        "pln": 0.166418,
                        "rub": 3.35,
                        "sar": 0.166057,
                        "sats": 120.67,
                        "sek": 0.374345,
                        "sgd": 0.05921,
                        "thb": 1.33,
                        "try": 0.315896,
                        "twd": 1.24,
                        "uah": 1.24,
                        "usd": 0.04426774,
                        "vef": 0.00443253,
                        "vnd": 1020.21,
                        "xag": 0.00169158,
                        "xau": 0.00002475,
                        "xdr": 0.03083151,
                        "xlm": 0.13717924,
                        "xrp": 0.1073201,
                        "yfi": 0.00000138,
                        "zar": 0.666482
                },
                "roi": null,
                "ath": {
                        "aed": 0.287137,
                        "ars": 6.81,
                        "aud": 0.102266,
                        "bch": 0.00019487,
                        "bdt": 6.63,
                        "bhd": 0.02946959,
                        "bits": 2.33,
                        "bmd": 0.078175,
                        "bnb": 0.02690809,
                        "brl": 0.425359,
                        "btc": 0.00000284,
                        "cad": 0.100593,
                        "chf": 0.069573,
                        "clp": 57.54,
                        "cny": 0.505277,
                        "czk": 1.69,
                        "dkk": 0.480635,
                        "dot": 0.00474765,
                        "eos": 0.02913307,
                        "eth": 0.00027029,
                        "eur": 0.06463,
                        "gbp": 0.057058,
                        "hkd": 0.606072,
                        "huf": 23.2,
                        "idr": 1099.73,
                        "ils": 0.2566,
                        "inr": 5.71,
                        "jpy": 8.17,
                        "krw": 87.3,
                        "kwd": 0.02366611,
                        "link": 0.00348451,
                        "lkr": 15.04,
                        "ltc": 0.00058245,
                        "mmk": 103.96,
                        "mxn": 1.59,
                        "myr": 0.316887,
                        "ngn": 29.8,
                        "nok": 0.676043,
                        "nzd": 0.109234,
                        "php": 3.76,
                        "pkr": 12.56,
                        "pln": 0.293497,
                        "rub": 5.96,
                        "sar": 0.293169,
                        "sats": 232.83,
                        "sek": 0.654299,
                        "sgd": 0.104086,
                        "thb": 2.34,
                        "try": 0.576329,
                        "twd": 2.19,
                        "uah": 2.2,
                        "usd": 0.078175,
                        "vef": 3360.53,
                        "vnd": 1803.72,
                        "xag": 0.00299121,
                        "xau": 0.00004244,
                        "xdr": 0.05429,
                        "xlm": 0.27087546,
                        "xrp": 0.29810806,
                        "yfi": 0.00000268,
                        "zar": 1.19
                },
                "ath_change_percentage": {
                        "aed": -43.08636,
                        "ars": -42.58619,
                        "aud": -42.71603,
                        "bch": -45.28238,
                        "bdt": -43.10845,
                        "bhd": -43.08228,
                        "bits": -47.74878,
                        "bmd": -43.08636,
                        "bnb": -96.90575,
                        "brl": -43.21606,
                        "btc": -57.19883,
                        "cad": -43.19778,
                        "chf": -42.23333,
                        "clp": -42.83585,
                        "cny": -43.01239,
                        "czk": -42.93381,
                        "dkk": -42.5121,
                        "dot": -52.42955,
                        "eos": -49.75092,
                        "eth": -89.59382,
                        "eur": -42.50733,
                        "gbp": -42.91294,
                        "hkd": -43.08783,
                        "huf": -42.96032,
                        "idr": -43.12055,
                        "ils": -42.92996,
                        "inr": -43.12089,
                        "jpy": -42.56478,
                        "krw": -42.89118,
                        "kwd": -43.03786,
                        "link": -45.39738,
                        "lkr": -42.79214,
                        "ltc": -47.01063,
                        "mmk": -40.62824,
                        "mxn": -42.69997,
                        "myr": -43.01686,
                        "ngn": -43.26552,
                        "nok": -43.08115,
                        "nzd": -43.07422,
                        "php": -43.05295,
                        "pkr": -43.2395,
                        "pln": -43.0114,
                        "rub": -43.49847,
                        "sar": -43.07382,
                        "sats": -47.74878,
                        "sek": -42.46085,
                        "sgd": -42.82613,
                        "thb": -42.9335,
                        "try": -44.91146,
                        "twd": -43.18033,
                        "uah": -43.55064,
                        "usd": -43.08636,
                        "vef": -99.99987,
                        "vnd": -43.13847,
                        "xag": -43.16247,
                        "xau": -41.42788,
                        "xdr": -42.9218,
                        "xlm": -48.99874,
                        "xrp": -63.59648,
                        "yfi": -47.96832,
                        "zar": -43.85373
                },
                "ath_date": {
                        "aed": "2021-01-29T04:08:39.575Z",
                        "ars": "2021-01-29T04:08:39.575Z",
                        "aud": "2021-01-29T04:08:39.575Z",
                        "bch": "2021-01-29T04:08:39.575Z",
                        "bdt": "2021-01-29T04:08:39.575Z",
                        "bhd": "2021-01-29T04:08:39.575Z",
                        "bits": "2021-01-29T04:08:39.575Z",
                        "bmd": "2021-01-29T04:08:39.575Z",
                        "bnb": "2017-10-19T00:00:00.000Z",
                        "brl": "2021-01-29T04:08:39.575Z",
                        "btc": "2014-02-12T00:00:00.000Z",
                        "cad": "2021-01-29T04:08:39.575Z",
                        "chf": "2021-01-29T04:08:39.575Z",
                        "clp": "2021-01-29T04:08:39.575Z",
                        "cny": "2021-01-29T04:08:39.575Z",
                        "czk": "2021-01-29T04:08:39.575Z",
                        "dkk": "2021-01-29T04:08:39.575Z",
                        "dot": "2021-01-29T04:08:39.575Z",
                        "eos": "2021-01-29T04:08:39.575Z",
                        "eth": "2015-10-20T00:00:00.000Z",
                        "eur": "2021-01-29T04:08:39.575Z",
                        "gbp": "2021-01-29T04:08:39.575Z",
                        "hkd": "2021-01-29T04:08:39.575Z",
                        "huf": "2021-01-29T04:08:39.575Z",
                        "idr": "2021-01-29T04:08:39.575Z",
                        "ils": "2021-01-29T04:08:39.575Z",
                        "inr": "2021-01-29T04:08:39.575Z",
                        "jpy": "2021-01-29T04:08:39.575Z",
                        "krw": "2021-01-29T04:08:39.575Z",
                        "kwd": "2021-01-29T04:08:39.575Z",
                        "link": "2021-01-29T04:08:39.575Z",
                        "lkr": "2021-01-29T04:08:39.575Z",
                        "ltc": "2021-01-29T04:08:39.575Z",
                        "mmk": "2021-01-29T04:08:39.575Z",
                        "mxn": "2021-01-29T04:08:39.575Z",
                        "myr": "2021-01-29T04:08:39.575Z",
                        "ngn": "2021-01-29T04:08:39.575Z",
                        "nok": "2021-01-29T04:08:39.575Z",
                        "nzd": "2021-01-29T04:08:39.575Z",
                        "php": "2021-01-29T04:08:39.575Z",
                        "pkr": "2021-01-29T04:08:39.575Z",
                        "pln": "2021-01-29T04:08:39.575Z",
                        "rub": "2021-01-29T04:08:39.575Z",
                        "sar": "2021-01-29T04:08:39.575Z",
                        "sats": "2021-01-29T04:08:39.575Z",
                        "sek": "2021-01-29T04:08:39.575Z",
                        "sgd": "2021-01-29T04:08:39.575Z",
                        "thb": "2021-01-29T04:08:39.575Z",
                        "try": "2021-01-29T04:08:39.575Z",
                        "twd": "2021-01-29T04:08:39.575Z",
                        "uah": "2021-01-29T04:08:39.575Z",
                        "usd": "2021-01-29T04:08:39.575Z",
                        "vef": "2021-01-03T12:14:20.368Z",
                        "vnd": "2021-01-29T04:08:39.575Z",
                        "xag": "2021-01-29T04:08:39.575Z",
                        "xau": "2021-01-29T04:08:39.575Z",
                        "xdr": "2021-01-29T04:08:39.575Z",
                        "xlm": "2021-01-29T04:08:39.575Z",
                        "xrp": "2021-01-29T04:08:39.575Z",
                        "yfi": "2021-01-29T04:08:39.575Z",
                        "zar": "2021-01-29T04:08:39.575Z"
                },
                "atl": {
                        "aed": 0.0003192,
                        "ars": 0.000776,
                        "aud": 0.00010901,
                        "bch": 6.819e-7,
                        "bdt": 0.00676005,
                        "bhd": 0.00003278,
                        "bits": 0.15856,
                        "bmd": 0.0000869,
                        "bnb": 0.00008194,
                        "brl": 0.00024014,
                        "btc": 1.509e-7,
                        "cad": 0.0001047,
                        "chf": 0.00007997,
                        "clp": 0.053147,
                        "cny": 0.00053661,
                        "czk": 0.00210666,
                        "dkk": 0.00057401,
                        "dot": 0.00045876,
                        "eos": 0.00039266,
                        "eth": 0.00000287,
                        "eur": 0.00007662,
                        "gbp": 0.00005701,
                        "hkd": 0.00067367,
                        "huf": 0.02345499,
                        "idr": 1.13,
                        "ils": 0.00033506,
                        "inr": 0.00552883,
                        "jpy": 0.01039259,
                        "krw": 0.09391,
                        "kwd": 0.0000262,
                        "link": 0.000178,
                        "lkr": 0.01159203,
                        "ltc": 0.00000856,
                        "mmk": 0.094336,
                        "mxn": 0.00133638,
                        "myr": 0.00031058,
                        "ngn": 0.948406,
                        "nok": 0.00064855,
                        "nzd": 0.00011599,
                        "php": 0.00387064,
                        "pkr": 0.00884262,
                        "pln": 0.00031075,
                        "rub": 0.00382859,
                        "sar": 0.00032591,
                        "sats": 15.86,
                        "sek": 0.0007174,
                        "sgd": 0.00011526,
                        "thb": 0.00289367,
                        "try": 0.00022968,
                        "twd": 0.00266709,
                        "uah": 0.00138586,
                        "usd": 0.0000869,
                        "vef": 0.0005491,
                        "vnd": 1.88,
                        "xag": 0.00000526,
                        "xau": 7e-8,
                        "xdr": 0.00006179,
                        "xlm": 0.01049371,
                        "xrp": 0.00501292,
                        "yfi": 6.47e-8,
                        "zar": 0.0010443
                },
                "atl_change_percentage": {
                        "aed": 51096.05803,
                        "ars": 503906.68924,
                        "aud": 53638.08526,
                        "bch": 15536.44693,
                        "bdt": 55686.3655,
                        "bhd": 51071.83234,
                        "bits": 667.27308,
                        "bmd": 51097.21493,
                        "bnb": 916.07115,
                        "brl": 100480.64811,
                        "btc": 706.02906,
                        "cad": 54472.85846,
                        "chf": 50155.4477,
                        "clp": 61786.07414,
                        "cny": 53560.03424,
                        "czk": 45564.1247,
                        "dkk": 48035.9278,
                        "dot": 392.29653,
                        "eos": 3628.22763,
                        "eth": 879.59319,
                        "eur": 48397.8771,
                        "gbp": 57033.35684,
                        "hkd": 51101.08388,
                        "huf": 56331.03801,
                        "idr": 55173.51688,
                        "ils": 43605.55169,
                        "inr": 58610.64728,
                        "jpy": 45038.23384,
                        "krw": 52991.36866,
                        "kwd": 51357.73894,
                        "link": 968.91382,
                        "lkr": 74131.61951,
                        "ltc": 3506.12028,
                        "mmk": 65330.69787,
                        "mxn": 68097.90759,
                        "myr": 58039.65112,
                        "ngn": 1682.68291,
                        "nok": 59231.48652,
                        "nzd": 53511.18521,
                        "php": 55205.79972,
                        "pkr": 80505.80521,
                        "pln": 53724.26139,
                        "rub": 87865.50438,
                        "sar": 51107.13967,
                        "sats": 667.27308,
                        "sek": 52378.47825,
                        "sgd": 51532.0278,
                        "thb": 46135.74381,
                        "try": 138130.72664,
                        "twd": 46603.72095,
                        "uah": 89571.5322,
                        "usd": 51097.21493,
                        "vef": 711.33255,
                        "vnd": 54428.56939,
                        "xag": 32213.36298,
                        "xau": 33996.0839,
                        "xdr": 50047.50619,
                        "xlm": 1216.50166,
                        "xrp": 2064.8438,
                        "yfi": 2055.4484,
                        "zar": 64107.05142
                },
                "atl_date": {
                        "aed": "2015-05-06T00:00:00.000Z",
                        "ars": "2015-05-06T00:00:00.000Z",
                        "aud": "2015-05-06T00:00:00.000Z",
                        "bch": "2017-11-12T00:00:00.000Z",
                        "bdt": "2015-05-06T00:00:00.000Z",
                        "bhd": "2015-05-06T00:00:00.000Z",
                        "bits": "2020-12-31T03:22:15.353Z",
                        "bmd": "2015-05-06T00:00:00.000Z",
                        "bnb": "2020-10-16T01:49:44.619Z",
                        "brl": "2014-08-18T00:00:00.000Z",
                        "btc": "2020-12-17T09:18:05.654Z",
                        "cad": "2015-05-06T00:00:00.000Z",
                        "chf": "2015-05-06T00:00:00.000Z",
                        "clp": "2015-05-06T00:00:00.000Z",
                        "cny": "2015-05-06T00:00:00.000Z",
                        "czk": "2015-05-06T00:00:00.000Z",
                        "dkk": "2015-05-06T00:00:00.000Z",
                        "dot": "2021-01-24T00:09:11.447Z",
                        "eos": "2019-05-28T00:00:00.000Z",
                        "eth": "2017-09-21T00:00:00.000Z",
                        "eur": "2015-05-06T00:00:00.000Z",
                        "gbp": "2015-05-06T00:00:00.000Z",
                        "hkd": "2015-05-06T00:00:00.000Z",
                        "huf": "2015-05-06T00:00:00.000Z",
                        "idr": "2015-05-06T00:00:00.000Z",
                        "ils": "2015-05-06T00:00:00.000Z",
                        "inr": "2015-05-06T00:00:00.000Z",
                        "jpy": "2015-05-06T00:00:00.000Z",
                        "krw": "2015-05-06T00:00:00.000Z",
                        "kwd": "2015-05-06T00:00:00.000Z",
                        "link": "2020-08-16T08:04:24.745Z",
                        "lkr": "2015-05-06T00:00:00.000Z",
                        "ltc": "2014-01-06T00:00:00.000Z",
                        "mmk": "2015-05-06T00:00:00.000Z",
                        "mxn": "2015-05-06T00:00:00.000Z",
                        "myr": "2015-05-06T00:00:00.000Z",
                        "ngn": "2020-11-03T04:09:14.796Z",
                        "nok": "2015-05-06T00:00:00.000Z",
                        "nzd": "2015-05-06T00:00:00.000Z",
                        "php": "2015-05-06T00:00:00.000Z",
                        "pkr": "2015-05-06T00:00:00.000Z",
                        "pln": "2015-05-06T00:00:00.000Z",
                        "rub": "2014-08-18T00:00:00.000Z",
                        "sar": "2015-05-06T00:00:00.000Z",
                        "sats": "2020-12-31T03:22:15.353Z",
                        "sek": "2015-05-06T00:00:00.000Z",
                        "sgd": "2015-05-06T00:00:00.000Z",
                        "thb": "2015-05-06T00:00:00.000Z",
                        "try": "2014-08-18T00:00:00.000Z",
                        "twd": "2015-05-06T00:00:00.000Z",
                        "uah": "2014-08-18T00:00:00.000Z",
                        "usd": "2015-05-06T00:00:00.000Z",
                        "vef": "2015-05-06T00:00:00.000Z",
                        "vnd": "2015-05-06T00:00:00.000Z",
                        "xag": "2015-05-06T00:00:00.000Z",
                        "xau": "2015-05-06T00:00:00.000Z",
                        "xdr": "2015-05-06T00:00:00.000Z",
                        "xlm": "2018-11-19T00:00:00.000Z",
                        "xrp": "2018-11-20T00:00:00.000Z",
                        "yfi": "2020-09-12T20:05:21.635Z",
                        "zar": "2015-05-06T00:00:00.000Z"
                },
                "market_cap": {
                        "aed": 20949206950,
                        "ars": 501353933967,
                        "aud": 7511417146,
                        "bch": 13628314,
                        "bdt": 483436681271,
                        "bhd": 2150222495,
                        "bits": 155701063595,
                        "bmd": 5703568459,
                        "bnb": 106490959,
                        "brl": 30931877645,
                        "btc": 155701,
                        "cad": 7322725991,
                        "chf": 5152113239,
                        "clp": 4216648161723,
                        "cny": 36912354353,
                        "czk": 123305121411,
                        "dkk": 35417386320,
                        "dot": 288382070,
                        "eos": 1870953531,
                        "eth": 3589179,
                        "eur": 4763095649,
                        "gbp": 4175645208,
                        "hkd": 44217028550,
                        "huf": 1696634948513,
                        "idr": 80192457711667,
                        "ils": 18772725226,
                        "inr": 416056240646,
                        "jpy": 601429886862,
                        "krw": 6392450733255,
                        "kwd": 1728124207,
                        "link": 242535587,
                        "lkr": 1103088231686,
                        "ltc": 39411981,
                        "mmk": 7912590660797,
                        "mxn": 116872602682,
                        "myr": 23147932591,
                        "ngn": 2167356014412,
                        "nok": 49328053281,
                        "nzd": 7970959361,
                        "php": 274495639225,
                        "pkr": 912856131860,
                        "pln": 21451212231,
                        "rub": 431595299216,
                        "sar": 21395151857,
                        "sats": 15570106359472,
                        "sek": 48251903985,
                        "sgd": 7629321313,
                        "thb": 171599996081,
                        "try": 40724476922,
                        "twd": 159674245090,
                        "uah": 159307169127,
                        "usd": 5703568459,
                        "vef": 571098310,
                        "vnd": 131446326193242,
                        "xag": 218324273,
                        "xau": 3190576,
                        "xdr": 3972409953,
                        "xlm": 17692814556,
                        "xrp": 13847080469,
                        "yfi": 177992,
                        "zar": 85945664327
                },
                "market_cap_rank": 14,
                "total_volume": {
                        "aed": 37224770546,
                        "ars": 890881086329,
                        "aud": 13343038098,
                        "bch": 24162836,
                        "bdt": 859021516996,
                        "bhd": 3820743152,
                        "bits": 276272864453,
                        "bmd": 10134704750,
                        "bnb": 188955435,
                        "brl": 54911350540,
                        "btc": 276273,
                        "cad": 13004802461,
                        "chf": 9153989641,
                        "clp": 7496641103450,
                        "cny": 65589782200,
                        "czk": 219091368257,
                        "dkk": 62926006808,
                        "dot": 510793390,
                        "eos": 3320074130,
                        "eth": 6368568,
                        "eur": 8462437927,
                        "gbp": 7417519463,
                        "hkd": 78569197226,
                        "huf": 3014669274885,
                        "idr": 142453916698858,
                        "ils": 33357367214,
                        "inr": 739331920738,
                        "jpy": 1068593134117,
                        "krw": 11354655534023,
                        "kwd": 3070714192,
                        "link": 428770143,
                        "lkr": 1960084045901,
                        "ltc": 69931833,
                        "mmk": 14059929451220,
                        "mxn": 207458662932,
                        "myr": 41131699227,
                        "ngn": 3851187804936,
                        "nok": 87558578992,
                        "nzd": 14154817814,
                        "php": 487722531381,
                        "pkr": 1622059495210,
                        "pln": 38099902301,
                        "rub": 766649875506,
                        "sar": 38017172706,
                        "sats": 27627286445314,
                        "sek": 85703026034,
                        "sgd": 13555704742,
                        "thb": 304851918875,
                        "try": 72321607809,
                        "twd": 283736261529,
                        "uah": 283073857226,
                        "usd": 10134704750,
                        "vef": 1014787987,
                        "vnd": 233567759552577,
                        "xag": 387271594,
                        "xau": 5665807,
                        "xdr": 7058598895,
                        "xlm": 31405966760,
                        "xrp": 24569983642,
                        "yfi": 315167,
                        "zar": 152585074302
                },
                "high_24h": {
                        "aed": 0.211652,
                        "ars": 5.06,
                        "aud": 0.075485,
                        "bch": 0.00013091,
                        "bdt": 4.88,
                        "bhd": 0.02172671,
                        "bits": 1.5,
                        "bmd": 0.057624,
                        "bnb": 0.00109515,
                        "brl": 0.308499,
                        "btc": 0.0000015,
                        "cad": 0.073649,
                        "chf": 0.051914,
                        "clp": 42.22,
                        "cny": 0.372317,
                        "czk": 1.24,
                        "dkk": 0.356973,
                        "dot": 0.00275929,
                        "eos": 0.01878541,
                        "eth": 0.00003444,
                        "eur": 0.04799041,
                        "gbp": 0.04241724,
                        "hkd": 0.446733,
                        "huf": 17.07,
                        "idr": 807.6,
                        "ils": 0.190385,
                        "inr": 4.2,
                        "jpy": 6.06,
                        "krw": 64.41,
                        "kwd": 0.01746108,
                        "link": 0.00234221,
                        "lkr": 11.13,
                        "ltc": 0.00037608,
                        "mmk": 79.94,
                        "mxn": 1.17,
                        "myr": 0.233721,
                        "ngn": 22.71,
                        "nok": 0.495825,
                        "nzd": 0.08002,
                        "php": 2.77,
                        "pkr": 9.24,
                        "pln": 0.215564,
                        "rub": 4.35,
                        "sar": 0.216159,
                        "sats": 150.46,
                        "sek": 0.485346,
                        "sgd": 0.076902,
                        "thb": 1.73,
                        "try": 0.412057,
                        "twd": 1.61,
                        "uah": 1.61,
                        "usd": 0.057624,
                        "vef": 0.00576984,
                        "vnd": 1327.7,
                        "xag": 0.00217448,
                        "xau": 0.0000316,
                        "xdr": 0.04010965,
                        "xlm": 0.17324836,
                        "xrp": 0.1476869,
                        "yfi": 0.00000171,
                        "zar": 0.861884
                },
                "low_24h": {
                        "aed": 0.114882,
                        "ars": 2.75,
                        "aud": 0.04104654,
                        "bch": 0.00007126,
                        "bdt": 2.65,
                        "bhd": 0.01179107,
                        "bits": 0.851302,
                        "bmd": 0.03127746,
                        "bnb": 0.00061483,
                        "brl": 0.167445,
                        "btc": 8.513e-7,
                        "cad": 0.03996478,
                        "chf": 0.02812075,
                        "clp": 22.9,
                        "cny": 0.202065,
                        "czk": 0.673812,
                        "dkk": 0.193471,
                        "dot": 0.00163686,
                        "eos": 0.01012269,
                        "eth": 0.00001957,
                        "eur": 0.02601306,
                        "gbp": 0.02290902,
                        "hkd": 0.24245,
                        "huf": 9.24,
                        "idr": 438.57,
                        "ils": 0.10325,
                        "inr": 2.28,
                        "jpy": 3.29,
                        "krw": 34.9,
                        "kwd": 0.0094756,
                        "link": 0.00125504,
                        "lkr": 6.06,
                        "ltc": 0.00020561,
                        "mmk": 42.08,
                        "mxn": 0.631654,
                        "myr": 0.126439,
                        "ngn": 11.92,
                        "nok": 0.268821,
                        "nzd": 0.04345612,
                        "php": 1.5,
                        "pkr": 5.01,
                        "pln": 0.116663,
                        "rub": 2.38,
                        "sar": 0.117325,
                        "sats": 85.13,
                        "sek": 0.263121,
                        "sgd": 0.04170443,
                        "thb": 0.938621,
                        "try": 0.223656,
                        "twd": 0.874299,
                        "uah": 0.875587,
                        "usd": 0.03127746,
                        "vef": 0.00313181,
                        "vnd": 719.58,
                        "xag": 0.00116099,
                        "xau": 0.00001704,
                        "xdr": 0.02175297,
                        "xlm": 0.09148522,
                        "xrp": 0.07716738,
                        "yfi": 9.287e-7,
                        "zar": 0.467877
                },
                "price_change_24h": 0.01299028,
                "price_change_percentage_24h": 41.53239,
                "price_change_percentage_7d": 495.20937,
                "price_change_percentage_14d": 389.35009,
                "price_change_percentage_30d": 361.51884,
                "price_change_percentage_60d": 1204.02481,
                "price_change_percentage_200d": 1165.79994,
                "price_change_percentage_1y": 1714.30717,
                "market_cap_change_24h": 1672287158,
                "market_cap_change_percentage_24h": 41.48277,
                "price_change_24h_in_currency": {
                        "aed": 0.04771329,
                        "ars": 1.14,
                        "aud": 0.01723499,
                        "bch": 0.00003378,
                        "bdt": 1.1,
                        "bhd": 0.00489769,
                        "bits": 0.355056,
                        "bmd": 0.01299028,
                        "bnb": 0.00021032,
                        "brl": 0.072404,
                        "btc": 3.551e-7,
                        "cad": 0.01683936,
                        "chf": 0.01186328,
                        "clp": 9.84,
                        "cny": 0.084427,
                        "czk": 0.283165,
                        "dkk": 0.081386,
                        "dot": 0.00059348,
                        "eos": 0.00414904,
                        "eth": 0.00000808,
                        "eur": 0.01095033,
                        "gbp": 0.00949023,
                        "hkd": 0.100735,
                        "huf": 3.93,
                        "idr": 183.66,
                        "ils": 0.0424528,
                        "inr": 0.951343,
                        "jpy": 1.38,
                        "krw": 14.69,
                        "kwd": 0.00393708,
                        "link": 0.0006178,
                        "lkr": 2.51,
                        "ltc": 0.00009795,
                        "mmk": 19.33,
                        "mxn": 0.274512,
                        "myr": 0.053221,
                        "ngn": 4.9,
                        "nok": 0.11363,
                        "nzd": 0.01837121,
                        "php": 0.627818,
                        "pkr": 2.07,
                        "pln": 0.04975443,
                        "rub": 0.969861,
                        "sar": 0.04873127,
                        "sats": 35.51,
                        "sek": 0.111224,
                        "sgd": 0.01750602,
                        "thb": 0.392953,
                        "try": 0.09224,
                        "twd": 0.365043,
                        "uah": 0.360862,
                        "usd": 0.01299028,
                        "vef": 0.00130072,
                        "vnd": 300.62,
                        "xag": 0.00052904,
                        "xau": 0.00000771,
                        "xdr": 0.00907853,
                        "xlm": 0.04549932,
                        "xrp": 0.02664041,
                        "yfi": 4.479e-7,
                        "zar": 0.198604
                },
                "price_change_percentage_1h_in_currency": {
                        "aed": -3.92067,
                        "ars": -3.92286,
                        "aud": -3.88432,
                        "bch": -4.96054,
                        "bdt": -3.92067,
                        "bhd": -3.85079,
                        "bits": -4.84769,
                        "bmd": -3.92067,
                        "bnb": -5.55557,
                        "brl": -3.69276,
                        "btc": -4.84769,
                        "cad": -3.90719,
                        "chf": -3.95481,
                        "clp": -3.63406,
                        "cny": -3.91919,
                        "czk": -3.9303,
                        "dkk": -3.95525,
                        "dot": -7.48772,
                        "eos": -6.14866,
                        "eth": -5.85939,
                        "eur": -3.95265,
                        "gbp": -3.92684,
                        "hkd": -3.92637,
                        "huf": -3.78002,
                        "idr": -3.80195,
                        "ils": -3.76805,
                        "inr": -3.90083,
                        "jpy": -3.95028,
                        "krw": -3.83622,
                        "kwd": -3.90957,
                        "link": -5.28041,
                        "lkr": -3.92067,
                        "ltc": -5.14875,
                        "mmk": -3.92067,
                        "mxn": -3.39227,
                        "myr": -3.92067,
                        "ngn": -4.22313,
                        "nok": -3.9652,
                        "nzd": -3.87504,
                        "php": -3.93893,
                        "pkr": -4.01064,
                        "pln": -3.86285,
                        "rub": -3.86679,
                        "sar": -3.90953,
                        "sats": -4.84769,
                        "sek": -3.9698,
                        "sgd": -3.93468,
                        "thb": -3.92067,
                        "try": -4.01785,
                        "twd": -3.86917,
                        "uah": -3.92067,
                        "usd": -3.92067,
                        "vef": -3.92067,
                        "vnd": -3.95357,
                        "xag": -3.895,
                        "xau": -3.95675,
                        "xdr": -3.92067,
                        "xlm": -5.33745,
                        "xrp": -7.92833,
                        "yfi": -6.13422,
                        "zar": -4.01375
                },
                "price_change_percentage_24h_in_currency": {
                        "aed": 41.53239,
                        "ars": 41.68922,
                        "aud": 41.98891,
                        "bch": 47.07575,
                        "bdt": 41.38108,
                        "bhd": 41.53727,
                        "bits": 41.68865,
                        "bmd": 41.53239,
                        "bnb": 34.19715,
                        "brl": 43.24023,
                        "btc": 41.68865,
                        "cad": 42.13551,
                        "chf": 42.18694,
                        "clp": 42.98197,
                        "cny": 41.78214,
                        "czk": 42.02435,
                        "dkk": 42.06631,
                        "dot": 36.23983,
                        "eos": 40.07638,
                        "eth": 40.9428,
                        "eur": 42.0955,
                        "gbp": 41.42572,
                        "hkd": 41.54882,
                        "huf": 42.46546,
                        "idr": 41.87652,
                        "ils": 41.11651,
                        "inr": 41.762,
                        "jpy": 42.05243,
                        "krw": 42.09906,
                        "kwd": 41.54967,
                        "link": 49.22536,
                        "lkr": 41.37506,
                        "ltc": 47.20233,
                        "mmk": 45.94917,
                        "mxn": 43.45924,
                        "myr": 42.09257,
                        "ngn": 41.08685,
                        "nok": 42.26965,
                        "nzd": 42.27531,
                        "php": 41.78425,
                        "pkr": 41.39987,
                        "pln": 42.64781,
                        "rub": 40.77072,
                        "sar": 41.53518,
                        "sats": 41.68865,
                        "sek": 42.27102,
                        "sgd": 41.9764,
                        "thb": 41.86488,
                        "try": 41.24213,
                        "twd": 41.75264,
                        "uah": 41.21371,
                        "usd": 41.53239,
                        "vef": 41.53239,
                        "vnd": 41.77755,
                        "xag": 45.50741,
                        "xau": 45.2237,
                        "xdr": 41.73467,
                        "xlm": 49.62845,
                        "xrp": 33.01998,
                        "yfi": 48.22983,
                        "zar": 42.44795
                },
                "price_change_percentage_7d_in_currency": {
                        "aed": 495.17697,
                        "ars": 500.64982,
                        "aud": 499.5561,
                        "bch": 433.59607,
                        "bdt": 494.85123,
                        "bhd": 495.35308,
                        "bits": 391.66613,
                        "bmd": 495.20937,
                        "bnb": 354.46878,
                        "brl": 495.72064,
                        "btc": 391.66613,
                        "cad": 496.03385,
                        "chf": 504.48513,
                        "clp": 496.01527,
                        "cny": 494.12611,
                        "czk": 497.06814,
                        "dkk": 501.32659,
                        "dot": 363.70176,
                        "eos": 387.05328,
                        "eth": 365.86156,
                        "eur": 501.469,
                        "gbp": 495.4624,
                        "hkd": 495.2556,
                        "huf": 494.09093,
                        "idr": 495.38672,
                        "ils": 499.85858,
                        "inr": 494.62137,
                        "jpy": 502.50071,
                        "krw": 501.5849,
                        "kwd": 495.70483,
                        "link": 429.14538,
                        "lkr": 497.0983,
                        "ltc": 403.14145,
                        "mmk": 520.96718,
                        "mxn": 499.32693,
                        "myr": 497.12205,
                        "ngn": 483.63503,
                        "nok": 493.42542,
                        "nzd": 493.91326,
                        "php": 495.82013,
                        "pkr": 492.69602,
                        "pln": 495.38058,
                        "rub": 493.15342,
                        "sar": 495.32015,
                        "sats": 391.66613,
                        "sek": 501.45235,
                        "sgd": 498.46804,
                        "thb": 496.65525,
                        "try": 474.20932,
                        "twd": 496.44073,
                        "uah": 489.90116,
                        "usd": 495.20937,
                        "vef": 495.20937,
                        "vnd": 493.8413,
                        "xag": 472.36203,
                        "xau": 512.31769,
                        "xdr": 497.98259,
                        "xlm": 341.29994,
                        "xrp": 262.78343,
                        "yfi": 414.79303,
                        "zar": 486.10663
                },
                "price_change_percentage_14d_in_currency": {
                        "aed": 389.32345,
                        "ars": 398.13016,
                        "aud": 399.76469,
                        "bch": 482.97723,
                        "bdt": 389.10067,
                        "bhd": 389.36437,
                        "bits": 374.91944,
                        "bmd": 389.35009,
                        "bnb": 286.94593,
                        "brl": 401.01515,
                        "btc": 374.91944,
                        "cad": 397.35304,
                        "chf": 397.05159,
                        "clp": 401.27726,
                        "cny": 389.75117,
                        "czk": 391.21749,
                        "dkk": 394.88665,
                        "dot": 342.8414,
                        "eos": 345.62274,
                        "eth": 324.78428,
                        "eur": 395.04581,
                        "gbp": 389.58957,
                        "hkd": 389.40312,
                        "huf": 393.73957,
                        "idr": 390.64538,
                        "ils": 392.54654,
                        "inr": 389.61431,
                        "jpy": 398.32032,
                        "krw": 399.59414,
                        "kwd": 389.69434,
                        "link": 353.63897,
                        "lkr": 384.10946,
                        "ltc": 405.96526,
                        "mmk": 409.29273,
                        "mxn": 411.42743,
                        "myr": 390.98327,
                        "ngn": 379.78751,
                        "nok": 398.86584,
                        "nzd": 391.11141,
                        "php": 390.60073,
                        "pkr": 387.80661,
                        "pln": 391.35624,
                        "rub": 403.71728,
                        "sar": 389.26309,
                        "sats": 374.91944,
                        "sek": 396.2567,
                        "sgd": 394.10252,
                        "thb": 391.73645,
                        "try": 370.87386,
                        "twd": 391.32968,
                        "uah": 383.74967,
                        "usd": 389.35009,
                        "vef": 389.35009,
                        "vnd": 388.90214,
                        "xag": 383.73089,
                        "xau": 411.93168,
                        "xdr": 391.16885,
                        "xlm": 347.08563,
                        "xrp": 251.26377,
                        "yfi": 419.68533,
                        "zar": 394.49918
                },
                "price_change_percentage_30d_in_currency": {
                        "aed": 361.5314,
                        "ars": 378.98492,
                        "aud": 366.10637,
                        "bch": 340.20821,
                        "bdt": 361.48808,
                        "bhd": 361.54332,
                        "bits": 293.61414,
                        "bmd": 361.51884,
                        "bnb": 242.55516,
                        "brl": 376.03767,
                        "btc": 293.61414,
                        "cad": 363.69495,
                        "chf": 373.0541,
                        "clp": 385.06037,
                        "cny": 362.24737,
                        "czk": 366.30165,
                        "dkk": 371.75031,
                        "dot": 117.75696,
                        "eos": 316.47844,
                        "eth": 195.76377,
                        "eur": 372.04637,
                        "gbp": 358.0734,
                        "hkd": 361.43985,
                        "huf": 365.73248,
                        "idr": 365.14043,
                        "ils": 373.38142,
                        "inr": 360.84109,
                        "jpy": 371.86561,
                        "krw": 376.21892,
                        "kwd": 359.84634,
                        "link": 163.35875,
                        "lkr": 377.54146,
                        "ltc": 383.77131,
                        "mmk": 382.35073,
                        "mxn": 375.9921,
                        "myr": 367.50885,
                        "ngn": 344.47461,
                        "nok": 366.49094,
                        "nzd": 362.78559,
                        "php": 362.48048,
                        "pkr": 362.08183,
                        "pln": 366.89138,
                        "rub": 371.83635,
                        "sar": 361.44515,
                        "sats": 293.61414,
                        "sek": 373.77157,
                        "sgd": 367.70489,
                        "thb": 363.98685,
                        "try": 343.84804,
                        "twd": 360.11367,
                        "uah": 353.41652,
                        "usd": 361.51884,
                        "vef": -99.99981,
                        "vnd": 359.18446,
                        "xag": 379.68339,
                        "xau": 400.90685,
                        "xdr": 363.16274,
                        "xlm": 118.29112,
                        "xrp": 161.32177,
                        "yfi": 229.09632,
                        "zar": 373.10705
                },
                "price_change_percentage_60d_in_currency": {
                        "aed": 1204.02481,
                        "ars": 1303.77077,
                        "aud": 1174.73503,
                        "bch": 794.66693,
                        "bdt": 1203.54745,
                        "bhd": 1203.83805,
                        "bits": 580.50766,
                        "bmd": 1204.02481,
                        "bnb": 622.63275,
                        "brl": 1270.10301,
                        "btc": 580.50766,
                        "cad": 1208.90568,
                        "chf": 1220.07502,
                        "clp": 1196.66123,
                        "cny": 1192.08582,
                        "czk": 1191.14397,
                        "dkk": 1218.69278,
                        "dot": 241.15736,
                        "eos": 1183.82731,
                        "eth": 388.06458,
                        "eur": 1220.02485,
                        "gbp": 1182.67531,
                        "hkd": 1204.38487,
                        "huf": 1210.01868,
                        "idr": 1195.60928,
                        "ils": 1211.87277,
                        "inr": 1189.03996,
                        "jpy": 1219.6887,
                        "krw": 1247.94298,
                        "kwd": 1198.74393,
                        "link": 626.95632,
                        "lkr": 1258.46175,
                        "ltc": 645.44217,
                        "mmk": 1267.38833,
                        "mxn": 1250.03412,
                        "myr": 1203.54303,
                        "ngn": 1202.31124,
                        "nok": 1180.21348,
                        "nzd": 1183.55383,
                        "php": 1204.5157,
                        "pkr": 1204.02481,
                        "pln": 1229.57763,
                        "rub": 1230.83647,
                        "sar": 1203.95146,
                        "sats": 580.50766,
                        "sek": 1204.01186,
                        "sgd": 1206.71434,
                        "thb": 1197.98358,
                        "try": 1092.9449,
                        "twd": 1193.85751,
                        "uah": 1187.39869,
                        "usd": 1204.02481,
                        "vef": -99.99947,
                        "vnd": 1194.42311,
                        "xag": 1105.63413,
                        "xau": 1240.86533,
                        "xdr": 1205.98619,
                        "xlm": 593.50564,
                        "xrp": 1746.42742,
                        "yfi": 1090.07595,
                        "zar": 1190.44385
                },
                "price_change_percentage_200d_in_currency": {
                        "aed": 1165.79994,
                        "ars": 1456.95728,
                        "aud": 1065.72547,
                        "bch": 577.465,
                        "bdt": 1165.31353,
                        "bhd": 1165.14891,
                        "bits": 216.40022,
                        "bmd": 1165.79994,
                        "bnb": 304.6081,
                        "brl": 1173.16426,
                        "btc": 216.40022,
                        "cad": 1096.07376,
                        "chf": 1118.07007,
                        "clp": 1088.5165,
                        "cny": 1071.59179,
                        "czk": 1073.96621,
                        "dkk": 1106.28444,
                        "eos": 937.88451,
                        "eth": 87.6214,
                        "eur": 1108.02704,
                        "gbp": 1064.16991,
                        "hkd": 1165.57303,
                        "huf": 1118.13281,
                        "idr": 1102.46461,
                        "ils": 1111.67686,
                        "inr": 1132.89439,
                        "jpy": 1147.27546,
                        "krw": 1078.24317,
                        "kwd": 1146.86183,
                        "link": 326.65442,
                        "lkr": 1217.03046,
                        "ltc": 270.87698,
                        "mmk": 1178.78353,
                        "mxn": 1049.15211,
                        "myr": 1105.21972,
                        "ngn": 1141.30059,
                        "nok": 1076.61936,
                        "nzd": 1059.47862,
                        "php": 1133.77597,
                        "pkr": 1110.40347,
                        "pln": 1114.57387,
                        "rub": 1232.18001,
                        "sar": 1165.98826,
                        "sats": 216.40022,
                        "sek": 1082.3932,
                        "sgd": 1117.75011,
                        "thb": 1102.68961,
                        "try": 1216.68277,
                        "twd": 1102.55075,
                        "uah": 1196.32899,
                        "usd": 1165.79994,
                        "vef": -99.99949,
                        "vnd": 1157.95915,
                        "xag": 835.92122,
                        "xau": 1180.85263,
                        "xdr": 1123.72123,
                        "xlm": 295.78775,
                        "xrp": 512.58449,
                        "yfi": -68.86583,
                        "zar": 1042.3787
                },
                "price_change_percentage_1y_in_currency": {
                        "aed": 1714.20838,
                        "ars": 2539.82218,
                        "aud": 1498.24521,
                        "bch": 1559.97022,
                        "bdt": 1710.60162,
                        "bhd": 1713.96073,
                        "bits": 359.26942,
                        "bmd": 1714.30717,
                        "bnb": 524.58949,
                        "brl": 2213.91106,
                        "btc": 359.26942,
                        "cad": 1651.18273,
                        "chf": 1596.73117,
                        "clp": 1604.18046,
                        "cny": 1572.24466,
                        "czk": 1621.31638,
                        "dkk": 1567.59919,
                        "eos": 2399.83858,
                        "eth": 116.439,
                        "eur": 1576.05319,
                        "gbp": 1625.84469,
                        "hkd": 1711.23256,
                        "huf": 1669.10709,
                        "idr": 1755.70323,
                        "ils": 1633.12203,
                        "inr": 1755.67908,
                        "jpy": 1661.15791,
                        "krw": 1607.6362,
                        "kwd": 1708.27937,
                        "link": 112.19613,
                        "lkr": 1833.81353,
                        "ltc": 772.34784,
                        "mmk": 1619.4851,
                        "mxn": 1872.39916,
                        "myr": 1688.96196,
                        "ngn": 1793.75727,
                        "nok": 1587.97132,
                        "nzd": 1537.1219,
                        "php": 1616.36953,
                        "pkr": 1779.72741,
                        "pln": 1655.33002,
                        "rub": 2054.81093,
                        "sar": 1713.96818,
                        "sats": 359.26942,
                        "sek": 1490.37814,
                        "sgd": 1671.96452,
                        "thb": 1655.93177,
                        "try": 2063.41901,
                        "twd": 1578.92683,
                        "uah": 1924.90185,
                        "usd": 1714.30717,
                        "vef": -99.99927,
                        "vnd": 1697.66776,
                        "xag": 1126.93654,
                        "xau": 1501.26363,
                        "xdr": 1639.46414,
                        "xlm": 260.50116,
                        "xrp": 1018.04956,
                        "zar": 1735.87819
                },
                "market_cap_change_24h_in_currency": {
                        "aed": 6142310730,
                        "ars": 147385263661,
                        "aud": 2219771871,
                        "bch": 4533244,
                        "bdt": 141378024410,
                        "bhd": 630497976,
                        "bits": 46747792627,
                        "bmd": 1672287158,
                        "bnb": 27925863,
                        "brl": 9361551629,
                        "btc": 46748,
                        "cad": 2170163952,
                        "chf": 1526197429,
                        "clp": 1262525224089,
                        "cny": 10868664634,
                        "czk": 36425966234,
                        "dkk": 10474236396,
                        "dot": 79483706,
                        "eos": 542942655,
                        "eth": 1077295,
                        "eur": 1409702517,
                        "gbp": 1222316433,
                        "hkd": 12968289509,
                        "huf": 505011416013,
                        "idr": 23634790438027,
                        "ils": 5465062522,
                        "inr": 122421125150,
                        "jpy": 177924973495,
                        "krw": 1892617528679,
                        "kwd": 506835443,
                        "link": 83622813,
                        "lkr": 322558020355,
                        "ltc": 13196908,
                        "mmk": 2489219145174,
                        "mxn": 35453777715,
                        "myr": 6851477930,
                        "ngn": 630631582335,
                        "nok": 14666250083,
                        "nzd": 2370828346,
                        "php": 80785691298,
                        "pkr": 267044867382,
                        "pln": 6410203381,
                        "rub": 125069972289,
                        "sar": 6273271473,
                        "sats": 4674779262720,
                        "sek": 14335045289,
                        "sgd": 2253398071,
                        "thb": 50594428144,
                        "try": 11881888751,
                        "twd": 46985416073,
                        "uah": 46454786997,
                        "usd": 1672287158,
                        "vef": 167446113,
                        "vnd": 38679387078153,
                        "xag": 68387175,
                        "xau": 993447,
                        "xdr": 1168718309,
                        "xlm": 5702680330,
                        "xrp": 3143770489,
                        "yfi": 59545,
                        "zar": 25624997523
                },
                "market_cap_change_percentage_24h_in_currency": {
                        "aed": 41.48277,
                        "ars": 41.63794,
                        "aud": 41.94861,
                        "bch": 49.84287,
                        "bdt": 41.33151,
                        "bhd": 41.48765,
                        "bits": 42.90628,
                        "bmd": 41.48277,
                        "bnb": 35.54487,
                        "brl": 43.40014,
                        "btc": 42.90628,
                        "cad": 42.11815,
                        "chf": 42.09136,
                        "clp": 42.73773,
                        "cny": 41.73243,
                        "czk": 41.92716,
                        "dkk": 41.99244,
                        "dot": 38.04898,
                        "eos": 40.8839,
                        "eth": 42.88796,
                        "eur": 42.03809,
                        "gbp": 41.38775,
                        "hkd": 41.5002,
                        "huf": 42.38011,
                        "idr": 41.78884,
                        "ils": 41.06704,
                        "inr": 41.69158,
                        "jpy": 42.01249,
                        "krw": 42.05973,
                        "kwd": 41.50005,
                        "link": 52.62183,
                        "lkr": 41.3255,
                        "ltc": 50.34091,
                        "mmk": 45.898,
                        "mxn": 43.54494,
                        "myr": 42.04275,
                        "ngn": 41.03739,
                        "nok": 42.31243,
                        "nzd": 42.33523,
                        "php": 41.70446,
                        "pkr": 41.3503,
                        "pln": 42.61817,
                        "rub": 40.80249,
                        "sar": 41.48473,
                        "sats": 42.90628,
                        "sek": 42.26525,
                        "sgd": 41.91649,
                        "thb": 41.81165,
                        "try": 41.19564,
                        "twd": 41.69483,
                        "uah": 41.16421,
                        "usd": 41.48277,
                        "vef": 41.48277,
                        "vnd": 41.69523,
                        "xag": 45.61058,
                        "xau": 45.2157,
                        "xdr": 41.68498,
                        "xlm": 47.56144,
                        "xrp": 29.37195,
                        "yfi": 50.27145,
                        "zar": 42.48129
                },
                "total_supply": null,
                "circulating_supply": 128233897815.778,
                "sparkline_7d": {
                        "price": [
                                0.012014061894495389,
                                0.014019818025748277,
                                0.016728439080892145,
                                0.017083749319049494,
                                0.019231635738889363,
                                0.019549387197711016,
                                0.023917672399617538,
                                0.021743438076546374,
                                0.02368391372336438,
                                0.02389899281738941,
                                0.026210158899095073,
                                0.03773867145063341,
                                0.04357580472050507,
                                0.03849070654524974,
                                0.04292983087294813,
                                0.07075463177531012,
                                0.06994932745308706,
                                0.05578337110844928,
                                0.05537689389078232,
                                0.05083651055430693,
                                0.04775244968224611,
                                0.03956391137955185,
                                0.046314385450373266,
                                0.04982266076718522,
                                0.05989689125397326,
                                0.05295732223610584,
                                0.05732295661501753,
                                0.053960915246588526,
                                0.05221961574791085,
                                0.04558286849298661,
                                0.04665267961577317,
                                0.04815271265004275,
                                0.051960672477495584,
                                0.052853069068323565,
                                0.049127706698414236,
                                0.04804660564468788,
                                0.04800734362430036,
                                0.043863306186895386,
                                0.039104420146548434,
                                0.034237062103633545,
                                0.03397179137121905,
                                0.0315556063157606,
                                0.031000625810634088,
                                0.03144878792345252,
                                0.03152040125064935,
                                0.03643900444866813,
                                0.03346539677358757,
                                0.033540625797366504,
                                0.03312883444688577,
                                0.034339901276463754,
                                0.03276831928665565,
                                0.03246610670689415,
                                0.026870076645473074,
                                0.028473107387876118,
                                0.026796526880705683,
                                0.024542661035623702,
                                0.024874162928831055,
                                0.026815539812156845,
                                0.03126710564868802,
                                0.02758457485802758,
                                0.029143759134113503,
                                0.029494581533758637,
                                0.029169170954256207,
                                0.027903534976453502,
                                0.027434384193486497,
                                0.02837097697385929,
                                0.02941302996878712,
                                0.02813298918596444,
                                0.027912187349678257,
                                0.027716614248953865,
                                0.02820240261208512,
                                0.028206208977149816,
                                0.030174856073869237,
                                0.028994583392387648,
                                0.029253932199316855,
                                0.028250657769142026,
                                0.02979136947153343,
                                0.032683868327315097,
                                0.04433779564999158,
                                0.04016863557960063,
                                0.0429112474862034,
                                0.03904277608326834,
                                0.037853898402699766,
                                0.03798879183692234,
                                0.0371371249455239,
                                0.038664579210803574,
                                0.03719147048598901,
                                0.037439910159559824,
                                0.04042275631695884,
                                0.040997212520228817,
                                0.03926531621526664,
                                0.038824622283619256,
                                0.03827282436460144,
                                0.03942460137568699,
                                0.038934908298337095,
                                0.03955907473687807,
                                0.04178719915436378,
                                0.039350930385316105,
                                0.03993379115342393,
                                0.03987624065251002,
                                0.04007075726712541,
                                0.03920922419584907,
                                0.03600801055039179,
                                0.0369934441541265,
                                0.036982379910365515,
                                0.03663078113334209,
                                0.03433622776659891,
                                0.03458624009599537,
                                0.030493648585242387,
                                0.03272250573494831,
                                0.031765613312264074,
                                0.03263672400519214,
                                0.03279948105059439,
                                0.03348427597432302,
                                0.03427535598221601,
                                0.03351019312031583,
                                0.03299726555765538,
                                0.032989639088717174,
                                0.033602100602940575,
                                0.03307709090397087,
                                0.033102401865735286,
                                0.03278647288070291,
                                0.031541493857021534,
                                0.030746445268372852,
                                0.0313488717173926,
                                0.031187067757163164,
                                0.030833994400747033,
                                0.03128770473085673,
                                0.030999311724169663,
                                0.030734530803458948,
                                0.032205636171577554,
                                0.03144668999349312,
                                0.0329617773878604,
                                0.03274624399832532,
                                0.032824441594133016,
                                0.033111467607235816,
                                0.032705248181624666,
                                0.03246659643140518,
                                0.03235895330482058,
                                0.03236161732437956,
                                0.032205174483588894,
                                0.03273972756390561,
                                0.03274120729496568,
                                0.03217557548214767,
                                0.03164982950973632,
                                0.03176144923329088,
                                0.03171217872995526,
                                0.03166596512680904,
                                0.0318022297989014,
                                0.03214688535725104,
                                0.032411877573333366,
                                0.03260854036758081,
                                0.03457372776396409,
                                0.037085856149185685,
                                0.036266427596595076,
                                0.03684086397198098,
                                0.036345419230412866,
                                0.036937289239059146,
                                0.03892505643889743,
                                0.0384870500654022,
                                0.04075592787774004,
                                0.03999115117861234,
                                0.040038435597358764,
                                0.05506816357087359,
                                0.04768908354436131,
                                0.047996623058211355,
                                0.051666342367793425
                        ]
                },
                "last_updated": "2021-02-04T16:27:26.362Z"
        },
        "community_data": {
                "facebook_likes": null,
                "twitter_followers": 375502,
                "reddit_average_posts_48h": 9.833,
                "reddit_average_comments_48h": 6099.833,
                "reddit_subscribers": 861357,
                "reddit_accounts_active_48h": 44901,
                "telegram_channel_user_count": null
        },
        "developer_data": {
                "forks": 1026,
                "stars": 3413,
                "subscribers": 352,
                "total_issues": 549,
                "closed_issues": 478,
                "pull_requests_merged": 823,
                "pull_request_contributors": 80,
                "commit_count_4_weeks": 0
        },
        "public_interest_stats": {
                "alexa_rank": 145564,
                "bing_matches": 0
        },
        "status_updates": [],
        "last_updated": "2021-02-04T16:27:26.362Z",
        "tickers": [
                {
                        "base": "DOGE",
                        "target": "USDT",
                        "market": {
                                "name": "Binance",
                                "identifier": "binance",
                                "has_trading_incentive": false
                        },
                        "last": 0.044358,
                        "converted_last": {
                                "btc": 0.00000121,
                                "eth": 0.00002787,
                                "usd": 0.04429537
                        },
                        "volume": 51091992266.58991,
                        "converted_volume": {
                                "btc": 61781,
                                "eth": 1424158,
                                "usd": 2263138492
                        },
                        "timestamp": "2021-02-04T16:24:55+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "USDT",
                        "market": {
                                "name": "OKEx",
                                "identifier": "okex",
                                "has_trading_incentive": false
                        },
                        "last": 0.04424,
                        "converted_last": {
                                "btc": 0.00000121,
                                "eth": 0.00002784,
                                "usd": 0.04424237
                        },
                        "volume": 19654841288.6677,
                        "converted_volume": {
                                "btc": 23738,
                                "eth": 547211,
                                "usd": 869576765
                        },
                        "timestamp": "2021-02-04T16:24:14+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "USDT",
                        "market": {
                                "name": "Huobi Global",
                                "identifier": "huobi",
                                "has_trading_incentive": false
                        },
                        "last": 0.04397,
                        "converted_last": {
                                "btc": 0.0000012,
                                "eth": 0.00002763,
                                "usd": 0.04397559
                        },
                        "volume": 13892369385.751112,
                        "converted_volume": {
                                "btc": 16652,
                                "eth": 383854,
                                "usd": 610925144
                        },
                        "timestamp": "2021-02-04T16:26:04+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "BTC",
                        "market": {
                                "name": "Binance",
                                "identifier": "binance",
                                "has_trading_incentive": false
                        },
                        "last": 0.00000122,
                        "converted_last": {
                                "btc": 0.00000122,
                                "eth": 0.00002812,
                                "usd": 0.04469062
                        },
                        "volume": 11578329840.991804,
                        "converted_volume": {
                                "btc": 14126,
                                "eth": 325619,
                                "usd": 517442693
                        },
                        "timestamp": "2021-02-04T16:24:56+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "USDT",
                        "market": {
                                "name": "XT",
                                "identifier": "xt",
                                "has_trading_incentive": false
                        },
                        "last": 0.044166,
                        "converted_last": {
                                "btc": 0.00000121,
                                "eth": 0.00002779,
                                "usd": 0.04416822
                        },
                        "volume": 5256698346,
                        "converted_volume": {
                                "btc": 6338,
                                "eth": 146107,
                                "usd": 232179013
                        },
                        "timestamp": "2021-02-04T16:23:07+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "USDT",
                        "market": {
                                "name": "Digifinex",
                                "identifier": "digifinex",
                                "has_trading_incentive": false
                        },
                        "last": 0.044042,
                        "converted_last": {
                                "btc": 0.00000121,
                                "eth": 0.00002791,
                                "usd": 0.04407512
                        },
                        "volume": 4210949232.71,
                        "converted_volume": {
                                "btc": 5075,
                                "eth": 117544,
                                "usd": 185598075
                        },
                        "timestamp": "2021-02-04T16:18:02+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "TRY",
                        "market": {
                                "name": "Paribu",
                                "identifier": "paribu",
                                "has_trading_incentive": false
                        },
                        "last": 0.3122,
                        "converted_last": {
                                "btc": 0.00000119,
                                "eth": 0.00002748,
                                "usd": 0.04372442
                        },
                        "volume": 4331826262.2945,
                        "converted_volume": {
                                "btc": 5165,
                                "eth": 119055,
                                "usd": 189406584
                        },
                        "timestamp": "2021-02-04T16:25:12+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "USD",
                        "market": {
                                "name": "Bitcoin.com Exchange",
                                "identifier": "bitcoin_com",
                                "has_trading_incentive": false
                        },
                        "last": 0.04394405,
                        "converted_last": {
                                "btc": 0.0000012,
                                "eth": 0.00002778,
                                "usd": 0.04394405
                        },
                        "volume": 3656315440,
                        "converted_volume": {
                                "btc": 4393,
                                "eth": 101572,
                                "usd": 160673309
                        },
                        "timestamp": "2021-02-04T16:19:20+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "XDG",
                        "target": "USD",
                        "market": {
                                "name": "Kraken",
                                "identifier": "kraken",
                                "has_trading_incentive": false
                        },
                        "last": 0.04418,
                        "converted_last": {
                                "btc": 0.00000121,
                                "eth": 0.0000278,
                                "usd": 0.04418
                        },
                        "volume": 2672391533.885742,
                        "converted_volume": {
                                "btc": 3223,
                                "eth": 74297,
                                "usd": 118066258
                        },
                        "timestamp": "2021-02-04T16:24:19+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "BUSD",
                        "market": {
                                "name": "Binance",
                                "identifier": "binance",
                                "has_trading_incentive": false
                        },
                        "last": 0.044482,
                        "converted_last": {
                                "btc": 0.00000121,
                                "eth": 0.00002796,
                                "usd": 0.0444368
                        },
                        "volume": 3285702035.6966414,
                        "converted_volume": {
                                "btc": 3986,
                                "eth": 91880,
                                "usd": 146006094
                        },
                        "timestamp": "2021-02-04T16:23:54+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "USD",
                        "market": {
                                "name": "Binance US",
                                "identifier": "binance_us",
                                "has_trading_incentive": false
                        },
                        "last": 0.044,
                        "converted_last": {
                                "btc": 0.0000012,
                                "eth": 0.00002766,
                                "usd": 0.044
                        },
                        "volume": 2514177096.7363634,
                        "converted_volume": {
                                "btc": 3016,
                                "eth": 69535,
                                "usd": 110623792
                        },
                        "timestamp": "2021-02-04T16:25:35+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "USDT",
                        "market": {
                                "name": "ZB",
                                "identifier": "zb",
                                "has_trading_incentive": false
                        },
                        "last": 0.0441669,
                        "converted_last": {
                                "btc": 0.00000121,
                                "eth": 0.0000278,
                                "usd": 0.04416912
                        },
                        "volume": 6528223666,
                        "converted_volume": {
                                "btc": 7872,
                                "eth": 181452,
                                "usd": 288345899
                        },
                        "timestamp": "2021-02-04T16:23:30+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "USDT",
                        "market": {
                                "name": "MXC",
                                "identifier": "mxc",
                                "has_trading_incentive": false
                        },
                        "last": 0.0441997,
                        "converted_last": {
                                "btc": 0.0000012,
                                "eth": 0.00002777,
                                "usd": 0.04413729
                        },
                        "volume": 2055272822.38,
                        "converted_volume": {
                                "btc": 2476,
                                "eth": 57085,
                                "usd": 90714171
                        },
                        "timestamp": "2021-02-04T16:24:43+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "BTC",
                        "market": {
                                "name": "Bitcoin.com Exchange",
                                "identifier": "bitcoin_com",
                                "has_trading_incentive": false
                        },
                        "last": 0.0000011967,
                        "converted_last": {
                                "btc": 0.0000012,
                                "eth": 0.00002767,
                                "usd": 0.04376494
                        },
                        "volume": 1760076130,
                        "converted_volume": {
                                "btc": 2106,
                                "eth": 48695,
                                "usd": 77029623
                        },
                        "timestamp": "2021-02-04T16:19:15+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "USDT",
                        "market": {
                                "name": "BiKi",
                                "identifier": "biki",
                                "has_trading_incentive": false
                        },
                        "last": 0.044045,
                        "converted_last": {
                                "btc": 0.0000012,
                                "eth": 0.00002768,
                                "usd": 0.04403286
                        },
                        "volume": 2054950129.6426926,
                        "converted_volume": {
                                "btc": 2467,
                                "eth": 56876,
                                "usd": 90485323
                        },
                        "timestamp": "2021-02-04T16:25:19+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "THB",
                        "market": {
                                "name": "Bitkub",
                                "identifier": "bitkub",
                                "has_trading_incentive": false
                        },
                        "last": 1.309,
                        "converted_last": {
                                "btc": 0.00000119,
                                "eth": 0.00002738,
                                "usd": 0.04350799
                        },
                        "volume": 1589050918.033516,
                        "converted_volume": {
                                "btc": 1887,
                                "eth": 43506,
                                "usd": 69136413
                        },
                        "timestamp": "2021-02-04T16:24:35+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "SXC",
                        "market": {
                                "name": "XT",
                                "identifier": "xt",
                                "has_trading_incentive": false
                        },
                        "last": 0.282604,
                        "converted_last": {
                                "btc": 0.00000121,
                                "eth": 0.00002786,
                                "usd": 0.04426497
                        },
                        "volume": 1019741694,
                        "converted_volume": {
                                "btc": 1232,
                                "eth": 28405,
                                "usd": 45138836
                        },
                        "timestamp": "2021-02-04T16:23:09+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "XDG",
                        "target": "EUR",
                        "market": {
                                "name": "Kraken",
                                "identifier": "kraken",
                                "has_trading_incentive": false
                        },
                        "last": 0.0369814,
                        "converted_last": {
                                "btc": 0.00000121,
                                "eth": 0.00002787,
                                "usd": 0.04428337
                        },
                        "volume": 1807548632.545387,
                        "converted_volume": {
                                "btc": 2185,
                                "eth": 50371,
                                "usd": 80044352
                        },
                        "timestamp": "2021-02-04T16:24:21+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "EUR",
                        "market": {
                                "name": "Binance",
                                "identifier": "binance",
                                "has_trading_incentive": false
                        },
                        "last": 0.0371,
                        "converted_last": {
                                "btc": 0.00000121,
                                "eth": 0.00002796,
                                "usd": 0.04442539
                        },
                        "volume": 1572875849.0890296,
                        "converted_volume": {
                                "btc": 1908,
                                "eth": 43972,
                                "usd": 69875626
                        },
                        "timestamp": "2021-02-04T16:24:13+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "USDT",
                        "market": {
                                "name": "WhiteBIT",
                                "identifier": "whitebit",
                                "has_trading_incentive": false
                        },
                        "last": 0.0445586,
                        "converted_last": {
                                "btc": 0.00000122,
                                "eth": 0.00002819,
                                "usd": 0.04459287
                        },
                        "volume": 109672290,
                        "converted_volume": {
                                "btc": 133.728,
                                "eth": 3092,
                                "usd": 4890602
                        },
                        "timestamp": "2021-02-04T16:19:46+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "USDT",
                        "market": {
                                "name": "Gate.io",
                                "identifier": "gate",
                                "has_trading_incentive": false
                        },
                        "last": 0.044235,
                        "converted_last": {
                                "btc": 0.00000121,
                                "eth": 0.00002784,
                                "usd": 0.04423722
                        },
                        "volume": 1212486007.3239458,
                        "converted_volume": {
                                "btc": 1464,
                                "eth": 33753,
                                "usd": 53637015
                        },
                        "timestamp": "2021-02-04T16:24:00+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "CNC",
                        "market": {
                                "name": "AEX",
                                "identifier": "aex",
                                "has_trading_incentive": false
                        },
                        "last": 0.28504,
                        "converted_last": {
                                "btc": 0.00000122,
                                "eth": 0.00002829,
                                "usd": 0.04482506
                        },
                        "volume": 688960710.8889803,
                        "converted_volume": {
                                "btc": 843.18,
                                "eth": 19494,
                                "usd": 30882705
                        },
                        "timestamp": "2021-02-04T16:21:26+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "USD",
                        "market": {
                                "name": "Coinsbit",
                                "identifier": "coinsbit",
                                "has_trading_incentive": false
                        },
                        "last": 0.0439269,
                        "converted_last": {
                                "btc": 0.0000012,
                                "eth": 0.00002782,
                                "usd": 0.0439269
                        },
                        "volume": 964744712.7716129,
                        "converted_volume": {
                                "btc": 1159,
                                "eth": 26839,
                                "usd": 42378245
                        },
                        "timestamp": "2021-02-04T16:18:54+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "XDG",
                        "target": "XBT",
                        "market": {
                                "name": "Kraken",
                                "identifier": "kraken",
                                "has_trading_incentive": false
                        },
                        "last": 0.00000121,
                        "converted_last": {
                                "btc": 0.00000121,
                                "eth": 0.00002789,
                                "usd": 0.0443243
                        },
                        "volume": 635068396.6921345,
                        "converted_volume": {
                                "btc": 768.433,
                                "eth": 17714,
                                "usd": 28148962
                        },
                        "timestamp": "2021-02-04T16:24:20+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "IDR",
                        "market": {
                                "name": "Indodax",
                                "identifier": "indodax",
                                "has_trading_incentive": false
                        },
                        "last": 618,
                        "converted_last": {
                                "btc": 0.0000012,
                                "eth": 0.00002763,
                                "usd": 0.04395432
                        },
                        "volume": 707755934.0283574,
                        "converted_volume": {
                                "btc": 848.271,
                                "eth": 19554,
                                "usd": 31108934
                        },
                        "timestamp": "2021-02-04T16:25:38+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "QC",
                        "market": {
                                "name": "ZB",
                                "identifier": "zb",
                                "has_trading_incentive": false
                        },
                        "last": 0.283395,
                        "converted_last": {
                                "btc": 0.00000122,
                                "eth": 0.00002804,
                                "usd": 0.04455527
                        },
                        "volume": 525009630,
                        "converted_volume": {
                                "btc": 638.574,
                                "eth": 14720,
                                "usd": 23391946
                        },
                        "timestamp": "2021-02-04T16:23:25+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "USDT",
                        "market": {
                                "name": "BW.com",
                                "identifier": "bw",
                                "has_trading_incentive": false
                        },
                        "last": 0.0446833,
                        "converted_last": {
                                "btc": 0.00000122,
                                "eth": 0.00002812,
                                "usd": 0.04468555
                        },
                        "volume": 2041430016,
                        "converted_volume": {
                                "btc": 2490,
                                "eth": 57405,
                                "usd": 91222416
                        },
                        "timestamp": "2021-02-04T16:23:37+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "USDT",
                        "market": {
                                "name": "Nominex",
                                "identifier": "nominex",
                                "has_trading_incentive": false
                        },
                        "last": 0.0439098,
                        "converted_last": {
                                "btc": 0.0000012,
                                "eth": 0.00002759,
                                "usd": 0.04389769
                        },
                        "volume": 231353421.885677,
                        "converted_volume": {
                                "btc": 276.928,
                                "eth": 6384,
                                "usd": 10155882
                        },
                        "timestamp": "2021-02-04T16:25:26+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "USDT",
                        "market": {
                                "name": "Binance US",
                                "identifier": "binance_us",
                                "has_trading_incentive": false
                        },
                        "last": 0.0440543,
                        "converted_last": {
                                "btc": 0.0000012,
                                "eth": 0.00002768,
                                "usd": 0.04404215
                        },
                        "volume": 535337119.8497627,
                        "converted_volume": {
                                "btc": 642.903,
                                "eth": 14820,
                                "usd": 23577400
                        },
                        "timestamp": "2021-02-04T16:25:34+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "BTC",
                        "market": {
                                "name": "Bittrex",
                                "identifier": "bittrex",
                                "has_trading_incentive": false
                        },
                        "last": 0.00000121,
                        "converted_last": {
                                "btc": 0.00000121,
                                "eth": 0.00002789,
                                "usd": 0.04432415
                        },
                        "volume": 602365678.3876691,
                        "converted_volume": {
                                "btc": 728.862,
                                "eth": 16802,
                                "usd": 26699349
                        },
                        "timestamp": "2021-02-04T16:23:51+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "BTC",
                        "market": {
                                "name": "Huobi Global",
                                "identifier": "huobi",
                                "has_trading_incentive": false
                        },
                        "last": 0.0000012007,
                        "converted_last": {
                                "btc": 0.0000012,
                                "eth": 0.00002768,
                                "usd": 0.04403367
                        },
                        "volume": 423525479.3028723,
                        "converted_volume": {
                                "btc": 508.527,
                                "eth": 11722,
                                "usd": 18649383
                        },
                        "timestamp": "2021-02-04T16:25:07+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "USDT",
                        "market": {
                                "name": "OKEx Korea",
                                "identifier": "okex_korea",
                                "has_trading_incentive": false
                        },
                        "last": 0.044268,
                        "converted_last": {
                                "btc": 0.00000121,
                                "eth": 0.00002786,
                                "usd": 0.04427023
                        },
                        "volume": 9727681.2843,
                        "converted_volume": {
                                "btc": 11.756174,
                                "eth": 271,
                                "usd": 430647
                        },
                        "timestamp": "2021-02-04T16:23:12+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "USDT",
                        "market": {
                                "name": "Bitforex",
                                "identifier": "bitforex",
                                "has_trading_incentive": false
                        },
                        "last": 0.04398311,
                        "converted_last": {
                                "btc": 0.0000012,
                                "eth": 0.00002764,
                                "usd": 0.04397098
                        },
                        "volume": 418698772.3504,
                        "converted_volume": {
                                "btc": 502.016,
                                "eth": 11572,
                                "usd": 18410597
                        },
                        "timestamp": "2021-02-04T16:25:56+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "INR",
                        "market": {
                                "name": "WazirX",
                                "identifier": "wazirx",
                                "has_trading_incentive": false
                        },
                        "last": 3.3499,
                        "converted_last": {
                                "btc": 0.00000125,
                                "eth": 0.00002899,
                                "usd": 0.0459226
                        },
                        "volume": 464662380,
                        "converted_volume": {
                                "btc": 582.584,
                                "eth": 13469,
                                "usd": 21338504
                        },
                        "timestamp": "2021-02-04T16:22:32+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "TRY",
                        "market": {
                                "name": "Binance",
                                "identifier": "binance",
                                "has_trading_incentive": false
                        },
                        "last": 0.3166,
                        "converted_last": {
                                "btc": 0.00000121,
                                "eth": 0.0000279,
                                "usd": 0.04434065
                        },
                        "volume": 485002475.80732787,
                        "converted_volume": {
                                "btc": 587.069,
                                "eth": 13533,
                                "usd": 21505325
                        },
                        "timestamp": "2021-02-04T16:24:53+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "USDT",
                        "market": {
                                "name": "CoinTiger",
                                "identifier": "cointiger",
                                "has_trading_incentive": false
                        },
                        "last": 0.0450897,
                        "converted_last": {
                                "btc": 0.00000124,
                                "eth": 0.00002866,
                                "usd": 0.04517799
                        },
                        "volume": 257879915.43223694,
                        "converted_volume": {
                                "btc": 319.119,
                                "eth": 7391,
                                "usd": 11650497
                        },
                        "timestamp": "2021-02-04T16:16:00+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "USDT",
                        "market": {
                                "name": "Float SV",
                                "identifier": "floatsv",
                                "has_trading_incentive": false
                        },
                        "last": 0.0443,
                        "converted_last": {
                                "btc": 0.00000121,
                                "eth": 0.00002796,
                                "usd": 0.04429714
                        },
                        "volume": 474324.3926,
                        "converted_volume": {
                                "btc": 0.57364861,
                                "eth": 13.262229,
                                "usd": 21011
                        },
                        "timestamp": "2021-02-04T16:22:59+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "ETH",
                        "market": {
                                "name": "Huobi Global",
                                "identifier": "huobi",
                                "has_trading_incentive": false
                        },
                        "last": 0.00002768,
                        "converted_last": {
                                "btc": 0.0000012,
                                "eth": 0.00002766,
                                "usd": 0.04395535
                        },
                        "volume": 120990809.40911002,
                        "converted_volume": {
                                "btc": 145.18,
                                "eth": 3347,
                                "usd": 5318193
                        },
                        "timestamp": "2021-02-04T16:24:54+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "USD",
                        "market": {
                                "name": "Bittrex",
                                "identifier": "bittrex",
                                "has_trading_incentive": false
                        },
                        "last": 0.04434,
                        "converted_last": {
                                "btc": 0.00000121,
                                "eth": 0.0000279,
                                "usd": 0.04434
                        },
                        "volume": 308938373.68490934,
                        "converted_volume": {
                                "btc": 373.949,
                                "eth": 8620,
                                "usd": 13698327
                        },
                        "timestamp": "2021-02-04T16:23:43+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "USDT",
                        "market": {
                                "name": "Bittrex",
                                "identifier": "bittrex",
                                "has_trading_incentive": false
                        },
                        "last": 0.04499956,
                        "converted_last": {
                                "btc": 0.00000123,
                                "eth": 0.00002832,
                                "usd": 0.04500182
                        },
                        "volume": 249828440.83407828,
                        "converted_volume": {
                                "btc": 306.914,
                                "eth": 7075,
                                "usd": 11242735
                        },
                        "timestamp": "2021-02-04T16:23:52+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "USDT",
                        "market": {
                                "name": "Hoo.com",
                                "identifier": "hoo",
                                "has_trading_incentive": false
                        },
                        "last": 0.0441016,
                        "converted_last": {
                                "btc": 0.00000121,
                                "eth": 0.00002795,
                                "usd": 0.04413476
                        },
                        "volume": 92467212.2104035,
                        "converted_volume": {
                                "btc": 111.592,
                                "eth": 2585,
                                "usd": 4081018
                        },
                        "timestamp": "2021-02-04T16:18:50+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "USDT",
                        "market": {
                                "name": "BitMart",
                                "identifier": "bitmart",
                                "has_trading_incentive": false
                        },
                        "last": 0.044094,
                        "converted_last": {
                                "btc": 0.0000012,
                                "eth": 0.00002783,
                                "usd": 0.04409115
                        },
                        "volume": 170527012,
                        "converted_volume": {
                                "btc": 205.277,
                                "eth": 4746,
                                "usd": 7518733
                        },
                        "timestamp": "2021-02-04T16:22:14+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "USDT",
                        "market": {
                                "name": "AEX",
                                "identifier": "aex",
                                "has_trading_incentive": false
                        },
                        "last": 0.045177,
                        "converted_last": {
                                "btc": 0.00000123,
                                "eth": 0.00002851,
                                "usd": 0.04517302
                        },
                        "volume": 50260439.31300599,
                        "converted_volume": {
                                "btc": 61.988,
                                "eth": 1433,
                                "usd": 2270416
                        },
                        "timestamp": "2021-02-04T16:21:28+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "USDT",
                        "market": {
                                "name": "CoinEx",
                                "identifier": "coinex",
                                "has_trading_incentive": false
                        },
                        "last": 0.04398931,
                        "converted_last": {
                                "btc": 0.0000012,
                                "eth": 0.00002764,
                                "usd": 0.0439272
                        },
                        "volume": 195200034.21235278,
                        "converted_volume": {
                                "btc": 234.076,
                                "eth": 5396,
                                "usd": 8574590
                        },
                        "timestamp": "2021-02-04T16:24:43+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "ETH",
                        "market": {
                                "name": "Bitcoin.com Exchange",
                                "identifier": "bitcoin_com",
                                "has_trading_incentive": false
                        },
                        "last": 0.0000279261,
                        "converted_last": {
                                "btc": 0.00000121,
                                "eth": 0.00002797,
                                "usd": 0.04431941
                        },
                        "volume": 53911540,
                        "converted_volume": {
                                "btc": 65.235,
                                "eth": 1508,
                                "usd": 2389328
                        },
                        "timestamp": "2021-02-04T16:21:08+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "USD",
                        "market": {
                                "name": "FTX",
                                "identifier": "ftx_spot",
                                "has_trading_incentive": false
                        },
                        "last": 0.044196,
                        "converted_last": {
                                "btc": 0.00000121,
                                "eth": 0.00002778,
                                "usd": 0.044196
                        },
                        "volume": 101142053.42571726,
                        "converted_volume": {
                                "btc": 121.889,
                                "eth": 2810,
                                "usd": 4470074
                        },
                        "timestamp": "2021-02-04T16:25:12+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "ETH",
                        "market": {
                                "name": "Bittrex",
                                "identifier": "bittrex",
                                "has_trading_incentive": false
                        },
                        "last": 0.0000274,
                        "converted_last": {
                                "btc": 0.00000119,
                                "eth": 0.00002738,
                                "usd": 0.04356022
                        },
                        "volume": 75143005.96024252,
                        "converted_volume": {
                                "btc": 89.254,
                                "eth": 2057,
                                "usd": 3273246
                        },
                        "timestamp": "2021-02-04T16:25:29+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "USD",
                        "market": {
                                "name": "FTX.us",
                                "identifier": "ftx_us",
                                "has_trading_incentive": false
                        },
                        "last": 0.04413,
                        "converted_last": {
                                "btc": 0.00000121,
                                "eth": 0.00002797,
                                "usd": 0.04413
                        },
                        "volume": 18387911.47829141,
                        "converted_volume": {
                                "btc": 22.207871,
                                "eth": 514.359,
                                "usd": 811459
                        },
                        "timestamp": "2021-02-04T16:17:43+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "ETH",
                        "market": {
                                "name": "Digifinex",
                                "identifier": "digifinex",
                                "has_trading_incentive": false
                        },
                        "last": 0.00002795,
                        "converted_last": {
                                "btc": 0.00000121,
                                "eth": 0.00002806,
                                "usd": 0.04427212
                        },
                        "volume": 96044922.87,
                        "converted_volume": {
                                "btc": 116.371,
                                "eth": 2695,
                                "usd": 4252112
                        },
                        "timestamp": "2021-02-04T16:17:58+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "BTC",
                        "market": {
                                "name": "XT",
                                "identifier": "xt",
                                "has_trading_incentive": false
                        },
                        "last": 0.000001207,
                        "converted_last": {
                                "btc": 0.00000121,
                                "eth": 0.00002782,
                                "usd": 0.04421426
                        },
                        "volume": 56893534,
                        "converted_volume": {
                                "btc": 68.67,
                                "eth": 1583,
                                "usd": 2515505
                        },
                        "timestamp": "2021-02-04T16:23:07+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "AUD",
                        "market": {
                                "name": "Binance",
                                "identifier": "binance",
                                "has_trading_incentive": false
                        },
                        "last": 0.05854,
                        "converted_last": {
                                "btc": 0.00000121,
                                "eth": 0.00002797,
                                "usd": 0.04445059
                        },
                        "volume": 64607313.44159549,
                        "converted_volume": {
                                "btc": 78.398,
                                "eth": 1807,
                                "usd": 2871833
                        },
                        "timestamp": "2021-02-04T16:24:26+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "USDT",
                        "market": {
                                "name": "FTX.us",
                                "identifier": "ftx_us",
                                "has_trading_incentive": false
                        },
                        "last": 0.0442325,
                        "converted_last": {
                                "btc": 0.00000121,
                                "eth": 0.00002803,
                                "usd": 0.04422764
                        },
                        "volume": 2585110.9903577687,
                        "converted_volume": {
                                "btc": 3.129057,
                                "eth": 72.472,
                                "usd": 114333
                        },
                        "timestamp": "2021-02-04T16:17:43+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "USDT",
                        "market": {
                                "name": "ZBG",
                                "identifier": "zbg",
                                "has_trading_incentive": false
                        },
                        "last": 0.044482,
                        "converted_last": {
                                "btc": 0.00000121,
                                "eth": 0.00002799,
                                "usd": 0.04448438
                        },
                        "volume": 44775538.4,
                        "converted_volume": {
                                "btc": 54.374,
                                "eth": 1253,
                                "usd": 1991812
                        },
                        "timestamp": "2021-02-04T16:24:08+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "USDT",
                        "market": {
                                "name": "Hotbit",
                                "identifier": "hotbit",
                                "has_trading_incentive": false
                        },
                        "last": 0.04447559,
                        "converted_last": {
                                "btc": 0.00000121,
                                "eth": 0.00002807,
                                "usd": 0.04447168
                        },
                        "volume": 25018359,
                        "converted_volume": {
                                "btc": 30.377169,
                                "eth": 702.292,
                                "usd": 1112608
                        },
                        "timestamp": "2021-02-04T16:21:25+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "BTC",
                        "market": {
                                "name": "Gate.io",
                                "identifier": "gate",
                                "has_trading_incentive": false
                        },
                        "last": 0.000001199,
                        "converted_last": {
                                "btc": 0.0000012,
                                "eth": 0.00002777,
                                "usd": 0.0438483
                        },
                        "volume": 52367809.36487152,
                        "converted_volume": {
                                "btc": 62.789,
                                "eth": 1454,
                                "usd": 2296240
                        },
                        "timestamp": "2021-02-04T16:18:55+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "BRL",
                        "market": {
                                "name": "NovaDAX",
                                "identifier": "novadax",
                                "has_trading_incentive": false
                        },
                        "last": 0.239479,
                        "converted_last": {
                                "btc": 0.0000012,
                                "eth": 0.00002776,
                                "usd": 0.04415784
                        },
                        "volume": 60648771,
                        "converted_volume": {
                                "btc": 73.026,
                                "eth": 1683,
                                "usd": 2678119
                        },
                        "timestamp": "2021-02-04T16:25:56+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "USDT",
                        "market": {
                                "name": "BitMax",
                                "identifier": "bitmax",
                                "has_trading_incentive": false
                        },
                        "last": 0.0442782,
                        "converted_last": {
                                "btc": 0.00000121,
                                "eth": 0.00002782,
                                "usd": 0.04426599
                        },
                        "volume": 97152518,
                        "converted_volume": {
                                "btc": 117.266,
                                "eth": 2703,
                                "usd": 4300553
                        },
                        "timestamp": "2021-02-04T16:25:19+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "USDT",
                        "market": {
                                "name": "FTX",
                                "identifier": "ftx_spot",
                                "has_trading_incentive": false
                        },
                        "last": 0.043835,
                        "converted_last": {
                                "btc": 0.00000119,
                                "eth": 0.00002755,
                                "usd": 0.04382291
                        },
                        "volume": 25799685.570617087,
                        "converted_volume": {
                                "btc": 30.829413,
                                "eth": 710.671,
                                "usd": 1130617
                        },
                        "timestamp": "2021-02-04T16:25:13+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "BTC",
                        "market": {
                                "name": "BitMax",
                                "identifier": "bitmax",
                                "has_trading_incentive": false
                        },
                        "last": 0.000001201,
                        "converted_last": {
                                "btc": 0.0000012,
                                "eth": 0.00002769,
                                "usd": 0.04404468
                        },
                        "volume": 57046244,
                        "converted_volume": {
                                "btc": 68.513,
                                "eth": 1579,
                                "usd": 2512583
                        },
                        "timestamp": "2021-02-04T16:25:19+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "BRL",
                        "market": {
                                "name": "Binance",
                                "identifier": "binance",
                                "has_trading_incentive": false
                        },
                        "last": 0.2416,
                        "converted_last": {
                                "btc": 0.00000122,
                                "eth": 0.00002803,
                                "usd": 0.04454893
                        },
                        "volume": 43938255.362665564,
                        "converted_volume": {
                                "btc": 53.435,
                                "eth": 1232,
                                "usd": 1957402
                        },
                        "timestamp": "2021-02-04T16:24:06+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "VET",
                        "market": {
                                "name": "Oceanex",
                                "identifier": "oceanex",
                                "has_trading_incentive": false
                        },
                        "last": 1.5288,
                        "converted_last": {
                                "btc": 0.0000012,
                                "eth": 0.00002772,
                                "usd": 0.04384493
                        },
                        "volume": 64731101,
                        "converted_volume": {
                                "btc": 77.605,
                                "eth": 1794,
                                "usd": 2838131
                        },
                        "timestamp": "2021-02-04T16:19:15+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "USD",
                        "market": {
                                "name": "P2PB2B",
                                "identifier": "p2pb2b",
                                "has_trading_incentive": false
                        },
                        "last": 0.04580141,
                        "converted_last": {
                                "btc": 0.00000125,
                                "eth": 0.00002903,
                                "usd": 0.04580141
                        },
                        "volume": 106181383,
                        "converted_volume": {
                                "btc": 132.632,
                                "eth": 3082,
                                "usd": 4863257
                        },
                        "timestamp": "2021-02-04T15:23:34+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "USDT",
                        "market": {
                                "name": "WazirX",
                                "identifier": "wazirx",
                                "has_trading_incentive": false
                        },
                        "last": 0.0449345,
                        "converted_last": {
                                "btc": 0.00000123,
                                "eth": 0.00002836,
                                "usd": 0.0449316
                        },
                        "volume": 71799434,
                        "converted_volume": {
                                "btc": 88.078,
                                "eth": 2036,
                                "usd": 3226063
                        },
                        "timestamp": "2021-02-04T16:22:33+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "USDT",
                        "market": {
                                "name": "P2PB2B",
                                "identifier": "p2pb2b",
                                "has_trading_incentive": false
                        },
                        "last": 0.04525597,
                        "converted_last": {
                                "btc": 0.00000125,
                                "eth": 0.00002904,
                                "usd": 0.04529094
                        },
                        "volume": 96334095,
                        "converted_volume": {
                                "btc": 120.062,
                                "eth": 2797,
                                "usd": 4363061
                        },
                        "timestamp": "2021-02-04T15:27:56+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "USDT",
                        "market": {
                                "name": "LBank",
                                "identifier": "lbank",
                                "has_trading_incentive": false
                        },
                        "last": 0.044505,
                        "converted_last": {
                                "btc": 0.00000121,
                                "eth": 0.00002797,
                                "usd": 0.04444216
                        },
                        "volume": 35182032.5794,
                        "converted_volume": {
                                "btc": 42.683454,
                                "eth": 983.927,
                                "usd": 1563565
                        },
                        "timestamp": "2021-02-04T16:24:32+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "USDT",
                        "market": {
                                "name": "Bitexlive",
                                "identifier": "bitexlive",
                                "has_trading_incentive": false
                        },
                        "last": 0.0436902,
                        "converted_last": {
                                "btc": 0.0000012,
                                "eth": 0.00002764,
                                "usd": 0.04372763
                        },
                        "volume": 15321296.7124,
                        "converted_volume": {
                                "btc": 18.317761,
                                "eth": 423.49,
                                "usd": 669964
                        },
                        "timestamp": "2021-02-04T16:20:24+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "USDT",
                        "market": {
                                "name": "BigONE",
                                "identifier": "bigone",
                                "has_trading_incentive": false
                        },
                        "last": 0.044992,
                        "converted_last": {
                                "btc": 0.00000123,
                                "eth": 0.00002846,
                                "usd": 0.04503055
                        },
                        "volume": 16328441,
                        "converted_volume": {
                                "btc": 20.103554,
                                "eth": 464.776,
                                "usd": 735279
                        },
                        "timestamp": "2021-02-04T16:20:38+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "BTC",
                        "market": {
                                "name": "WhiteBIT",
                                "identifier": "whitebit",
                                "has_trading_incentive": false
                        },
                        "last": 0.00000121,
                        "converted_last": {
                                "btc": 0.00000121,
                                "eth": 0.00002797,
                                "usd": 0.04425134
                        },
                        "volume": 56423949,
                        "converted_volume": {
                                "btc": 68.273,
                                "eth": 1578,
                                "usd": 2496835
                        },
                        "timestamp": "2021-02-04T16:19:47+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "USDT",
                        "market": {
                                "name": "NovaDAX",
                                "identifier": "novadax",
                                "has_trading_incentive": false
                        },
                        "last": 0.044025,
                        "converted_last": {
                                "btc": 0.0000012,
                                "eth": 0.00002767,
                                "usd": 0.04401286
                        },
                        "volume": 7743846.04,
                        "converted_volume": {
                                "btc": 9.293641,
                                "eth": 214.235,
                                "usd": 340829
                        },
                        "timestamp": "2021-02-04T16:25:56+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "HUSD",
                        "market": {
                                "name": "Huobi Global",
                                "identifier": "huobi",
                                "has_trading_incentive": false
                        },
                        "last": 0.04382,
                        "converted_last": {
                                "btc": 0.00000119,
                                "eth": 0.00002767,
                                "usd": 0.04365315
                        },
                        "volume": 37147991.568620935,
                        "converted_volume": {
                                "btc": 44.380429,
                                "eth": 1028,
                                "usd": 1621627
                        },
                        "timestamp": "2021-02-04T16:17:51+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "USDT",
                        "market": {
                                "name": "Atomars",
                                "identifier": "atomars",
                                "has_trading_incentive": false
                        },
                        "last": 0.044224,
                        "converted_last": {
                                "btc": 0.00000121,
                                "eth": 0.00002783,
                                "usd": 0.04422622
                        },
                        "volume": 7834987.25890859,
                        "converted_volume": {
                                "btc": 9.459389,
                                "eth": 218.055,
                                "usd": 346512
                        },
                        "timestamp": "2021-02-04T16:23:05+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "BTC",
                        "market": {
                                "name": "Bitexlive",
                                "identifier": "bitexlive",
                                "has_trading_incentive": false
                        },
                        "last": 0.00000122,
                        "converted_last": {
                                "btc": 0.00000122,
                                "eth": 0.00002821,
                                "usd": 0.04462096
                        },
                        "volume": 9365136.609,
                        "converted_volume": {
                                "btc": 11.425467,
                                "eth": 264.146,
                                "usd": 417881
                        },
                        "timestamp": "2021-02-04T16:20:40+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "INR",
                        "market": {
                                "name": "BitBNS",
                                "identifier": "bitbns",
                                "has_trading_incentive": false
                        },
                        "last": 3.4,
                        "converted_last": {
                                "btc": 0.00000127,
                                "eth": 0.00002951,
                                "usd": 0.04660295
                        },
                        "volume": 13158538,
                        "converted_volume": {
                                "btc": 16.768239,
                                "eth": 388.371,
                                "usd": 613227
                        },
                        "timestamp": "2021-02-04T16:19:00+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "KRW",
                        "market": {
                                "name": "ProBit",
                                "identifier": "probit",
                                "has_trading_incentive": false
                        },
                        "last": 47.045,
                        "converted_last": {
                                "btc": 0.00000115,
                                "eth": 0.00002653,
                                "usd": 0.0419752
                        },
                        "volume": 11171108.1828,
                        "converted_volume": {
                                "btc": 12.820645,
                                "eth": 296.402,
                                "usd": 468909
                        },
                        "timestamp": "2021-02-04T16:20:17+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "BTC",
                        "market": {
                                "name": "NovaDAX",
                                "identifier": "novadax",
                                "has_trading_incentive": false
                        },
                        "last": 0.0000011965,
                        "converted_last": {
                                "btc": 0.0000012,
                                "eth": 0.00002771,
                                "usd": 0.04375688
                        },
                        "volume": 1537741.91,
                        "converted_volume": {
                                "btc": 1.839908,
                                "eth": 42.614294,
                                "usd": 67287
                        },
                        "timestamp": "2021-02-04T16:18:58+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "BTC",
                        "market": {
                                "name": "ProBit",
                                "identifier": "probit",
                                "has_trading_incentive": false
                        },
                        "last": 0.0000012099,
                        "converted_last": {
                                "btc": 0.00000121,
                                "eth": 0.00002797,
                                "usd": 0.04425156
                        },
                        "volume": 9518431.35021629,
                        "converted_volume": {
                                "btc": 11.51635,
                                "eth": 266.247,
                                "usd": 421205
                        },
                        "timestamp": "2021-02-04T16:20:23+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "THB",
                        "market": {
                                "name": "Satang Pro",
                                "identifier": "tdax",
                                "has_trading_incentive": false
                        },
                        "last": 1.45,
                        "converted_last": {
                                "btc": 0.00000132,
                                "eth": 0.00003062,
                                "usd": 0.04821215
                        },
                        "volume": 6957390.7645,
                        "converted_volume": {
                                "btc": 9.189549,
                                "eth": 213.036,
                                "usd": 335431
                        },
                        "timestamp": "2021-02-04T16:14:26+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "BTC",
                        "market": {
                                "name": "TradeOgre",
                                "identifier": "trade_ogre",
                                "has_trading_incentive": false
                        },
                        "last": 0.00000121,
                        "converted_last": {
                                "btc": 0.00000121,
                                "eth": 0.00002789,
                                "usd": 0.04437474
                        },
                        "volume": 4329895.586776859,
                        "converted_volume": {
                                "btc": 5.239174,
                                "eth": 120.772,
                                "usd": 192138
                        },
                        "timestamp": "2021-02-04T16:25:30+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "USDT",
                        "market": {
                                "name": "ProBit",
                                "identifier": "probit",
                                "has_trading_incentive": false
                        },
                        "last": 0.044651,
                        "converted_last": {
                                "btc": 0.00000122,
                                "eth": 0.00002825,
                                "usd": 0.04468926
                        },
                        "volume": 10234823.07709206,
                        "converted_volume": {
                                "btc": 12.505595,
                                "eth": 289.118,
                                "usd": 457387
                        },
                        "timestamp": "2021-02-04T16:20:21+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "USDT",
                        "market": {
                                "name": "LATOKEN",
                                "identifier": "latoken",
                                "has_trading_incentive": false
                        },
                        "last": 0.0449552,
                        "converted_last": {
                                "btc": 0.00000123,
                                "eth": 0.00002844,
                                "usd": 0.04499372
                        },
                        "volume": 2128613.0942404885,
                        "converted_volume": {
                                "btc": 2.618602,
                                "eth": 60.54,
                                "usd": 95774
                        },
                        "timestamp": "2021-02-04T16:20:28+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "BTC",
                        "market": {
                                "name": "CREX24",
                                "identifier": "crex24",
                                "has_trading_incentive": false
                        },
                        "last": 0.000001399,
                        "converted_last": {
                                "btc": 0.0000014,
                                "eth": 0.00003232,
                                "usd": 0.051151
                        },
                        "volume": 1626136.45378501,
                        "converted_volume": {
                                "btc": 2.274965,
                                "eth": 52.558,
                                "usd": 83179
                        },
                        "timestamp": "2021-02-04T15:45:37+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "EUR",
                        "market": {
                                "name": "Dove Wallet",
                                "identifier": "dove_wallet",
                                "has_trading_incentive": false
                        },
                        "last": 0.0369355,
                        "converted_last": {
                                "btc": 0.00000121,
                                "eth": 0.00002783,
                                "usd": 0.04422841
                        },
                        "volume": 184629.593425119,
                        "converted_volume": {
                                "btc": 0.22291852,
                                "eth": 5.138658,
                                "usd": 8165.87
                        },
                        "timestamp": "2021-02-04T16:24:23+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "BTC",
                        "market": {
                                "name": "SouthXchange",
                                "identifier": "south_xchange",
                                "has_trading_incentive": false
                        },
                        "last": 0.00000125,
                        "converted_last": {
                                "btc": 0.00000125,
                                "eth": 0.0000289,
                                "usd": 0.04578416
                        },
                        "volume": 811929.55173838,
                        "converted_volume": {
                                "btc": 1.014912,
                                "eth": 23.463832,
                                "usd": 37174
                        },
                        "timestamp": "2021-02-04T16:22:33+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "BTC",
                        "market": {
                                "name": "StakeCube Exchange",
                                "identifier": "stake_cube",
                                "has_trading_incentive": false
                        },
                        "last": 0.00000123,
                        "converted_last": {
                                "btc": 0.00000123,
                                "eth": 0.00002844,
                                "usd": 0.04498671
                        },
                        "volume": 1045365.081300813,
                        "converted_volume": {
                                "btc": 1.285799,
                                "eth": 29.726493,
                                "usd": 47028
                        },
                        "timestamp": "2021-02-04T16:20:58+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "KRW",
                        "market": {
                                "name": "Dove Wallet",
                                "identifier": "dove_wallet",
                                "has_trading_incentive": false
                        },
                        "last": 47.5038,
                        "converted_last": {
                                "btc": 0.00000116,
                                "eth": 0.00002687,
                                "usd": 0.0423912
                        },
                        "volume": 173192.693470409,
                        "converted_volume": {
                                "btc": 0.20093052,
                                "eth": 4.653772,
                                "usd": 7341.85
                        },
                        "timestamp": "2021-02-04T16:17:58+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "BTC",
                        "market": {
                                "name": "Graviex",
                                "identifier": "graviex",
                                "has_trading_incentive": false
                        },
                        "last": 0.000001315,
                        "converted_last": {
                                "btc": 0.00000132,
                                "eth": 0.00003049,
                                "usd": 0.04819116
                        },
                        "volume": 232279.3024,
                        "converted_volume": {
                                "btc": 0.30544728,
                                "eth": 7.081188,
                                "usd": 11193.81
                        },
                        "timestamp": "2021-02-04T16:06:54+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "USDT",
                        "market": {
                                "name": "Bitsten",
                                "identifier": "bitsten",
                                "has_trading_incentive": false
                        },
                        "last": 0.04633,
                        "converted_last": {
                                "btc": 0.00000127,
                                "eth": 0.00002931,
                                "usd": 0.0463697
                        },
                        "volume": 197689.02270839628,
                        "converted_volume": {
                                "btc": 0.25063267,
                                "eth": 5.794397,
                                "usd": 9166.78
                        },
                        "timestamp": "2021-02-04T16:20:12+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "BTC",
                        "market": {
                                "name": "qTrade",
                                "identifier": "qtrade",
                                "has_trading_incentive": false
                        },
                        "last": 0.00000126,
                        "converted_last": {
                                "btc": 0.00000126,
                                "eth": 0.00002905,
                                "usd": 0.04615588
                        },
                        "volume": 323331.17786399,
                        "converted_volume": {
                                "btc": 0.40739728,
                                "eth": 9.391212,
                                "usd": 14923.64
                        },
                        "timestamp": "2021-02-04T16:24:06+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "BTC",
                        "market": {
                                "name": "Dove Wallet",
                                "identifier": "dove_wallet",
                                "has_trading_incentive": false
                        },
                        "last": 0.00000122,
                        "converted_last": {
                                "btc": 0.00000122,
                                "eth": 0.00002812,
                                "usd": 0.04469062
                        },
                        "volume": 152198.528816938,
                        "converted_volume": {
                                "btc": 0.18568221,
                                "eth": 4.280296,
                                "usd": 6801.85
                        },
                        "timestamp": "2021-02-04T16:24:37+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "LTC",
                        "market": {
                                "name": "STEX",
                                "identifier": "stocks_exchange",
                                "has_trading_incentive": false
                        },
                        "last": 0.00030882,
                        "converted_last": {
                                "btc": 0.00000122,
                                "eth": 0.00002811,
                                "usd": 0.04446824
                        },
                        "volume": 20217.58010629,
                        "converted_volume": {
                                "btc": 0.02458103,
                                "eth": 0.56829075,
                                "usd": 899.04
                        },
                        "timestamp": "2021-02-04T16:20:51+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "DASH",
                        "market": {
                                "name": "SouthXchange",
                                "identifier": "south_xchange",
                                "has_trading_incentive": false
                        },
                        "last": 0.00040096,
                        "converted_last": {
                                "btc": 0.00000123,
                                "eth": 0.00002839,
                                "usd": 0.04497183
                        },
                        "volume": 6358.693077,
                        "converted_volume": {
                                "btc": 0.00780734,
                                "eth": 0.18049857,
                                "usd": 285.96
                        },
                        "timestamp": "2021-02-04T16:22:36+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "PPC",
                        "target": "DOGE",
                        "market": {
                                "name": "SouthXchange",
                                "identifier": "south_xchange",
                                "has_trading_incentive": false
                        },
                        "last": 8.99547267,
                        "converted_last": {
                                "btc": 0.00000128,
                                "eth": 0.00002965,
                                "usd": 0.04673065
                        },
                        "volume": 133.171847,
                        "converted_volume": {
                                "btc": 0.00153337,
                                "eth": 0.03551449,
                                "usd": 55.98
                        },
                        "timestamp": "2021-02-04T16:15:07+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "peercoin"
                },
                {
                        "base": "DOGE",
                        "target": "USDT",
                        "market": {
                                "name": "Changelly PRO",
                                "identifier": "changelly",
                                "has_trading_incentive": false
                        },
                        "last": 0.0435446,
                        "converted_last": {
                                "btc": 0.00000119,
                                "eth": 0.00002755,
                                "usd": 0.04357809
                        },
                        "volume": 3656243760,
                        "converted_volume": {
                                "btc": 4357,
                                "eth": 100724,
                                "usd": 159332110
                        },
                        "timestamp": "2021-02-04T16:19:01+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "TRY",
                        "market": {
                                "name": "Thodex",
                                "identifier": "thodex",
                                "has_trading_incentive": false
                        },
                        "last": 0.313,
                        "converted_last": {
                                "btc": 0.0000012,
                                "eth": 0.00002755,
                                "usd": 0.04383646
                        },
                        "volume": 3125841942,
                        "converted_volume": {
                                "btc": 3736,
                                "eth": 86130,
                                "usd": 137025847
                        },
                        "timestamp": "2021-02-04T16:25:32+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "BTC",
                        "market": {
                                "name": "Changelly PRO",
                                "identifier": "changelly",
                                "has_trading_incentive": false
                        },
                        "last": 0.00000118885,
                        "converted_last": {
                                "btc": 0.00000119,
                                "eth": 0.00002749,
                                "usd": 0.04347785
                        },
                        "volume": 1759135650,
                        "converted_volume": {
                                "btc": 2091,
                                "eth": 48350,
                                "usd": 76483442
                        },
                        "timestamp": "2021-02-04T16:19:02+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "BTC",
                        "market": {
                                "name": "VCC Exchange",
                                "identifier": "vcc",
                                "has_trading_incentive": false
                        },
                        "last": 0.00000121,
                        "converted_last": {
                                "btc": 0.00000121,
                                "eth": 0.00002797,
                                "usd": 0.04425134
                        },
                        "volume": 1648538040,
                        "converted_volume": {
                                "btc": 1995,
                                "eth": 46116,
                                "usd": 72950013
                        },
                        "timestamp": "2021-02-04T16:19:47+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "ETH",
                        "market": {
                                "name": "Changelly PRO",
                                "identifier": "changelly",
                                "has_trading_incentive": false
                        },
                        "last": 0.0000275152,
                        "converted_last": {
                                "btc": 0.00000119,
                                "eth": 0.00002758,
                                "usd": 0.04362172
                        },
                        "volume": 53914930,
                        "converted_volume": {
                                "btc": 64.309,
                                "eth": 1487,
                                "usd": 2351862
                        },
                        "timestamp": "2021-02-04T16:19:02+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "USDT",
                        "market": {
                                "name": "Bibox",
                                "identifier": "bibox",
                                "has_trading_incentive": false
                        },
                        "last": 0.044017,
                        "converted_last": {
                                "btc": 0.0000012,
                                "eth": 0.00002766,
                                "usd": 0.04400486
                        },
                        "volume": 67664650,
                        "converted_volume": {
                                "btc": 81.192,
                                "eth": 1872,
                                "usd": 2977574
                        },
                        "timestamp": "2021-02-04T16:25:17+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOG",
                        "target": "USD",
                        "market": {
                                "name": "Bitfinex",
                                "identifier": "bitfinex",
                                "has_trading_incentive": false
                        },
                        "last": 0.045399,
                        "converted_last": {
                                "btc": 0.00000124,
                                "eth": 0.00002857,
                                "usd": 0.045399
                        },
                        "volume": 97216934.38,
                        "converted_volume": {
                                "btc": 120.485,
                                "eth": 2777,
                                "usd": 4413552
                        },
                        "timestamp": "2021-02-04T16:24:50+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                },
                {
                        "base": "DOGE",
                        "target": "USDT",
                        "market": {
                                "name": "BiONE",
                                "identifier": "bione",
                                "has_trading_incentive": false
                        },
                        "last": 0.04425464,
                        "converted_last": {
                                "btc": 0.00000121,
                                "eth": 0.000028,
                                "usd": 0.04428867
                        },
                        "volume": 49644059.69151495,
                        "converted_volume": {
                                "btc": 60.12,
                                "eth": 1390,
                                "usd": 2198670
                        },
                        "timestamp": "2021-02-04T16:19:39+00:00",
                        "is_anomaly": false,
                        "is_stale": false,
                        "coin_id": "dogecoin"
                }
        ]
}`
