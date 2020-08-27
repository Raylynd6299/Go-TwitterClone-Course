import React, {useState} from 'react'
import { Form, Button, Spinner } from "react-bootstrap"
import { values, size } from "lodash";
import { toast } from "react-toastify"
import {isEmailValid} from "../../utils/validations"
import { signInApi, setTokenApi } from "../../api/auth"
import "./SignInForm.scss"


export default function SignInForm(props) {
    const {setRefreshCheckLogin} = props
    const [formData, setFormData] = useState(initialFormVale())
    const [signInLoading, setSignInLoading] = useState(false)

    const onSubmit = e =>{
        e.preventDefault();
        let validCount = 0
        values(formData).some(value => {
             value && validCount++
             return null
        })
        if (validCount !== size(formData)) {
            toast.warning("Completa todos los campos del formulario")
        }else{
            if (!isEmailValid(formData.email)) {
                toast.warning("Email invalido")
            }else if(size(formData.password) < 6){
                toast.warning("La contraseña tiene que tener al menos 6 caracteres")
            }else{
                setSignInLoading(true)
                signInApi(formData).then(response => {
                    if(response.message){
                        toast.warning(response.message)
                    }else{
                        setTokenApi(response.token)
                        setRefreshCheckLogin(true)
                    }
                }).catch(()=>{
                    toast.error("Error en el servidor intentelo mas tarde")
                }).finally(()=>{
                    setSignInLoading(false)
                })
            }
        }
    }
    const onChange = e =>{
        setFormData({
            ...formData,
            [e.target.name]:e.target.value
        });
    };

    return (
        <div className="sign-in-form">
            <h2>Log In</h2>
            <Form onSubmit={onSubmit} onChange={onChange}>
                <Form.Group>
                    <Form.Control type="email"  name="email" placeholder="Email" defaultValue={formData.email}/>
                </Form.Group>
                <Form.Group>
                    <Form.Control type="password" name="password" placeholder="Contraseña" defaultValue={formData.password}/>
                </Form.Group>
                <Button variant="primary" type="submit">
                    {!signInLoading ? "Log in": <Spinner animation ="border"/>}
                </Button>
            </Form>
        </div>
    )
}

function initialFormVale(){
    return {
        email:"",
        password:""
    }
}