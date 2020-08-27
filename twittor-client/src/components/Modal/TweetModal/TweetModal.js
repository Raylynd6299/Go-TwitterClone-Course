import React,{useState} from 'react';
import {Modal,Form,Button} from "react-bootstrap";
import {toast} from "react-toastify"
import  classNames from "classnames"
import {Close} from "../../../utils/Icons"
import {addTweetApi} from "../../../api/tweet"

import "./TweetModal.scss"

export default function TwettModal(props) {
    const {show,setShow} = props;
    const [message, setMessage] = useState("")
    const maxLength = 280;
    const onSubmit = (e) => {
        e.preventDefault();

        if (message.length > 0 && message.length <= maxLength){
            addTweetApi(message).then(response => {
                if(response?.code >= 200 && response?.code <300) {
                    toast.success(response.message)
                    setShow(false);
                    window.location.reload();
                }
            }).catch(()=> {
                toast.warning("Error al enviar el tweet,intentelo mas tarde")
            })
        }
    }

    return (
        <Modal className="tweet-modal" show={show} onHide={()=> setShow(false)} centered size="lg">
            <Modal.Header>
                <Modal.Title>
                    <Close onClick={() => setShow(false)}/>
                </Modal.Title>
            </Modal.Header>
            <Modal.Body>
                <Form onSubmit={onSubmit}>
                    <Form.Control as="textarea" rows="6" placeholder="Que esta pensando" onChange={(e) => setMessage(e.target.value)}/>
                    <span className={classNames("count",{error: message.length > maxLength})}>
                    {message.length}
                    </span>
                    <Button type="submit" disabled={message.length > maxLength || message.length <1}>Twittoar</Button>
                </Form>
            </Modal.Body>
        </Modal>
    )
}
