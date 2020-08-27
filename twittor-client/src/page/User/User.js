import React, { useState, useEffect } from 'react'
//import { Button, Spinner } from "react-bootstrap"
import BasicLayout from "../../layout/BasicLayout"
import { withRouter } from "react-router-dom"
import { getUserApi } from "../../api/user"
import BannerAvatar from "../../components/User/BannerAvatar"
import { toast } from "react-toastify"
import userAuth from "../../hooks/useAuth"
import InfoUser from "../../components/User/InfoUser"
import {getUserTweetsApi} from "../../api/tweet"
import ListTweets from "../../components/ListTweets"
import "./User.scss"
import useAuth from '../../hooks/useAuth'

function User(props) {
    const {match} = props;
    const {params} = match;
    const [user, setUser] = useState(null);
    const [tweets, setTweets] = useState(null);
    const loggedUser = userAuth();

    useEffect(() => {
        getUserApi(match.params.id).then(response => {
            setUser(response);
            if(!response) toast.error("El usuario que has visitado no existe")
        }).catch(() => {
            toast.error("El usuario que has visitado no existe")
        })
    }, [params])

    useEffect(() => {
        getUserTweetsApi(params.id,1).then((response) => {
            setTweets(response);
        }).catch(() => {
            setTweets([]);
        })
    }, [params])
    return (
        <BasicLayout className = "user">
            <div className = "user__title">
                <h2>{user?`${user.nombre} ${user.apellidos}`:"El usuario no existe"}</h2>
            </div>
            <BannerAvatar user={user} loggedUser={loggedUser}/>
            <InfoUser user={user}/>
            <div className = "user__tweets"> 
                <h3>Tweets</h3>
                {tweets && <ListTweets tweets={tweets}/>}
            </div>
        </BasicLayout>
    )
}
export default withRouter(User);
