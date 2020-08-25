import React from 'react'
import {API_HOST} from "../../../utils/constant"
import Avatar_NF from "../../../assets/png/avatar-no-found.png"
import { Button} from "react-bootstrap"

import "./BannerAvatar.scss"

export default function BannerAvatar(props) {
    const {user,loggedUser} = props;
    const bannerUrl = user?.banner ? `${API_HOST}/obtenerBanner?id=${user.id}` : null;
    const avatarURL = user?.avatar ? `${API_HOST}/obtenerAvatar?id=${user.id}` : Avatar_NF;

    return (
        <div className="banner-avatar" style={{backgroundImage:`url('${bannerUrl}')`}}>
            <div className="avatar" style={{backgroundImage:`url('${avatarURL}')`}}></div>
        
            {user && (
                <div className="options">
                {loggedUser._id === user.id && <Button>Editar Perfil</Button> }
                {loggedUser._id !== user.id && <Button>Seguir</Button> }
                </div>
            )}
        </div>
    )
}
