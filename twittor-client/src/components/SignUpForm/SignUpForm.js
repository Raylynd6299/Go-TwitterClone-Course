import React, {useState} from 'react'
import { Button, Form, Col, Row} from "react-bootstrap";

import "./SignUpForm.scss"

export default function SignUpForm(props) {
    const {setShowModal} = props;
    const [formData, setFormData] = useState(initialFormValue())

    const onSubmit = e => {
        e.preventDefault();
        setShowModal(false);
        console.log(formData)
    }

    const onChange = e => {
        setFormData({...formData,[e.target.name]:e.target.value})
    }
    //onChange={e=>setFormData({...formData,nombre: e.target.value})}

    return (
        <div>
            <h2>Crea tu cuenta</h2>
            <Form onSubmit = {onSubmit} onChange={onChange}>
                <Form.Group>
                    <Row>
                        <Col>
                            <Form.Control type="text" placeholder = "Nombre" name="nombre" defaultValue={formData.nombre} />
                        </Col>
                        <Col>
                            <Form.Control type="text" placeholder = "Apellidos" name="apellidos" defaultValue={formData.apellidos} />
                        </Col>
                    </Row>
                </Form.Group>
                <Form.Group>
                    <Form.Control type="email" placeholder = "Email" name="email" defaultValue={formData.email} />
                </Form.Group>
                <Form.Group>
                    <Row>
                        <Col>
                            <Form.Control type="password" placeholder = "Contraseña" name ="password" defaultValue={formData.password} />
                        </Col>
                        <Col>
                            <Form.Control type="password" placeholder = "Repetir contraseña" name="repeatPassword" defaultValue={formData.repeatPassword}/>
                        </Col>
                    </Row>
                </Form.Group>


                <Button variant="primary" type= "submit">Registrate</Button>
            </Form> 
        </div>
    )
}

function initialFormValue(){
    return {
        nombre: "",
        apellidos: "",
        email:"",
        password:"",
        repeatPassword:""
    };
}
