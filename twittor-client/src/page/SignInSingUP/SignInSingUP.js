import React, { useState } from 'react'
import { Container, Row, Col, Button } from "react-bootstrap"
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome"
import { faSearch, faUser, faUsers, faComment } from "@fortawesome/free-solid-svg-icons"
import {} from "@fortawesome/fontawesome-svg-core";

import BasicModal from "../../components/Modal/BasicModal"
import SignUpForm from "../../components/SignUpForm"

import LogoWhiteTwittor from "../../assets/png/logo-white.png"
import LogoTwittor from "../../assets/png/logo.png"

import "./SignInSingUP.scss"

export default function SignInSingUP() {
    const [showModal, setShowModal] = useState(false);
    const [contentModal, setContentModal] = useState(null);

    const openModal = content => {
        setShowModal(true);
        setContentModal(content);
    }

    return (
        <>
        <Container className="signin-signup" fluid>
            <Row>
                <LeftComponent/>
                <RightComponent openModal = {openModal} setShowModal = {setShowModal}/>
            </Row>
        </Container>
       <BasicModal show = {showModal} setShow = {setShowModal} >
           {contentModal}
       </BasicModal>
        </>
    )
}

function LeftComponent(){
    return (
        <Col className="signin-signup__left" xs={6}>
            <img src={LogoTwittor} alt="Twittor"></img>
            <div>
                <h2>
                    <FontAwesomeIcon icon={faSearch}/>
                     Sigue lo que te interesa
                </h2>
                <h2>
                    <FontAwesomeIcon icon={faUsers}/>
                     Enterate de que esta hablando la gente
                </h2>
                <h2>
                    <FontAwesomeIcon icon={faComment}/>
                    Unete a la conversacion
                </h2>
            </div>
        </Col>
    )
}
function RightComponent(props){
    const {openModal, setShowModal } = props;
    return (
        <Col className="signin-signup__right" xs={6}>
            <div>
                <img src={LogoWhiteTwittor} alt="Twittor"></img>
                <h2>Mira lo que esta pasando en el mundo en este momento</h2>
                <h3>Unete a Twittor hoy mismo.</h3>
                <Button onClick = {() => openModal(
                    <SignUpForm setShowModal = {setShowModal} />
                )} variant="primary">Regístrate</Button>
                <Button onClick = {() => openModal(
                    <h2>Formulario de Inicio de Login</h2>
                )} variant="outline-primary">Iniciar Sesion</Button>
            </div>
        </Col>
    )
}
