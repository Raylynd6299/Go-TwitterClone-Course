import React, {useState,useEffect} from 'react';
import BasicLayout from "../../layout/BasicLayout"
import {getTweetsFollowersApi} from "../../api/tweet"
import ListTweets from "../../components/ListTweets"

import "./Home.scss"

export default function Home(props) {
    const {setRefreshCheckLogin} =props;
    const [tweets, setTweets] = useState(null)
    const [page, setPage] = useState(1)

    useEffect(() => {
        getTweetsFollowersApi(page).then(response => {
            setTweets(formatModel(response))
        })
    }, [page])

    return (
        <BasicLayout className = "home" setRefreshCheckLogin = {setRefreshCheckLogin}>
            <div className="home__title">
                <h2>Inicio</h2>
            </div>
            {tweets && <ListTweets tweets={tweets}/>}
            <p>Cargar mas Tweets</p>
        </BasicLayout>
    )
}

function formatModel(tweets){
    const tweetsTemp = [];
    
    tweets.forEach(tweet => {
        tweetsTemp.push({
            _id:tweet._id,
            userid:tweet.usuarioRelationId,
            mensaje:tweet.Tweet.mensaje,
            fecha:tweet.Tweet.fecha
        })
    });
    return tweetsTemp
}