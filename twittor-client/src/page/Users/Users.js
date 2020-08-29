import React, {useState,useEffect} from 'react'
import BasicLayout from "../../layout/BasicLayout"
import {Spinner,ButtonGroup,Button} from "react-bootstrap";
import {getFolowsApi} from "../../api/follow"
import {withRouter} from "react-router-dom"
import queryString from "query-string"
import ListUsers from "../../components/ListUsers"
import { isEmpty } from "lodash";
import {useDebouncedCallback} from "use-debounce"

import "./Users.scss";


function Users(props) {
    const {setRefreshCheckLogin, location,history} = props;
    const [users, setUsers] = useState(null);
    const params = useUserQuery(location);
    const [typeUser, setTypeUser] = useState(params.type || "follow")
    const [btnloading, setBtnloading] = useState(false)

    const [onSearch] = useDebouncedCallback(value => {
        setUsers(null);
        history.push({search:queryString.stringify({...params,search:value, page:1})})
    },250)

    useEffect(() => {
        getFolowsApi(queryString.stringify(params)).then(response => {
            if(params.page == 1){
                if(isEmpty(response)){
                    setUsers([]);
                }else{
                    setUsers(response)
                }
            }else {
                if(!response){
                    setBtnloading(0)
                } else {
                    setUsers([...users,...response]);
                    setBtnloading(false)
                }
            }
            
        }).catch(() => {
            setUsers([]);
        });
    }, [location])

    const onChangeType = (type) => {
        setUsers(null);
        if(type === "new") {
            setBtnloading(false)
            setTypeUser("new")
        }else {
            setBtnloading(false)
            setTypeUser("follow")
        }
        history.push({
            search:queryString.stringify({type:type,page:1,search:""})
        })

    }
     
    const moreData = () => {
        setBtnloading(true);
        const newPage = parseInt( params.page) + 1;
        history.push({
            search:queryString.stringify({...params,page:newPage})
        })
    }

    return (
        <BasicLayout className="users" title="Usuarios" setRefreshCheckLogin={setRefreshCheckLogin}>
                <div className="users__title">
                    <h2> Usuarios</h2>
                    <input type="text" placeholder="Busca un usuario" onChange={(e)=> onSearch(e.target.value) }></input>
                </div>
                <ButtonGroup className="users__options">
                    <Button className={typeUser==="follow" && "active"} onClick={() => onChangeType("follow")}>Siguiendo</Button>
                    <Button className={typeUser==="new" && "active"} onClick={() => onChangeType("new")}>Nuevos</Button>
                </ButtonGroup>
                {!users ? (
                    <div className="users__loading">
                        <Spinner animation="border" variant="info"/>
                        Buscando usuarios
                    </div>
                ) : (
                    <>
                        <ListUsers users={users}/>
                        <Button onClick={moreData} className="load-more">
                            {!btnloading ? (
                                btnloading !== 0 && "Cargar mas usuarios"
                            ) : ( 
                                <Spinner as="span" animation="grow" size="sm" role="status" aria-hidden="true"/>
                            )}
                        </Button>
                    </>
                )}
        </BasicLayout>
    )
}

function useUserQuery(location){
    const {page=1,type="follow",search} =queryString.parse(location.search);
    return {page,type,search};
}

export default withRouter(Users);