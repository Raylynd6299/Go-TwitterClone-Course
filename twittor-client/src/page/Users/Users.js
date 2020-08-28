import React, {useState,useEffect} from 'react'
import BasicLayout from "../../layout/BasicLayout"
import {Spinner,ButtonGroup,Button} from "react-bootstrap";
import {getFolowsApi} from "../../api/follow"
import {withRouter} from "react-router-dom"
import queryString from "query-string"
import ListUsers from "../../components/ListUsers"
import { isEmpty } from "lodash";


import "./Users.scss";


function Users(props) {
    const {setRefreshCheckLogin, location} = props;
    const [users, setUsers] = useState(null);
    const params = useUserQuery(location);


    useEffect(() => {
        getFolowsApi(queryString.stringify(params)).then(response => {
            if(isEmpty(response)){
                setUsers([]);
            }else{
                setUsers(response)
            }
        }).catch(() => {
            setUsers([]);
        });
    }, [])
    return (
        <BasicLayout classNames="users" title="Usuarios" setRefreshCheckLogin={setRefreshCheckLogin}>
                <div className="users__title">
                    <h2> Usuarios</h2>
                    <input type="text" placeholder="Busca un usuario"></input>
                </div>
                <ButtonGroup className="users__options">
                    <Button className="active">Siguiendo</Button>
                    <Button className>Nuevos</Button>
                </ButtonGroup>
                {!users ? (
                    <div className="users__loading">
                        <Spinner animation="border" variant="info"/>
                        Buscando usuarios
                    </div>
                ) : (
                    <ListUsers users={users}/>
                )}

        </BasicLayout>
    )
}

function useUserQuery(location){
    const {page=1,type="follow",search} =queryString.parse(location.search);
    return {page,type,search};
}

export default withRouter(Users);