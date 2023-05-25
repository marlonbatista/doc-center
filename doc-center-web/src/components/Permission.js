import { borderRadius } from "@mui/system";
import React from "react";

const MyComponent = () => {
    const containerStyle = {
        maxWidth: '250px',
        margin: '0 auto',
        padding: '10%',
        background: 'none',
        borderRadius: '40px',
    };

    const divCard = {
        backgroundColor: '#000750',
        borderRadius: '10px',
        // maxWidth: '300px',
        padding: '5px',
        color: 'white',
    };

    const icons = {
        backgroundColor: 'white',
        borderRadius: '10px',
    };

    return (
        <div class="container">
            <div class="row">
                <div class="col card text-center" style={{backgroundColor: '#eff1f3'}}>
                    <div class="shadow" style={{marginLeft: '15px', marginTop: '25px', paddingTop: '20px', backgroundColor: '#000750', borderRadius: '10px', padding: '5px', color: 'white', fontWeight: 'bold', }}>
                        <h5 style={{marginTop: '5px'}}>Solicitação de acesso</h5>
                    </div>
                    <div class="row img-group" style={{}}>
                        <img style={containerStyle} class="img img-fluid bordered-img" src="https://img.estadao.com.br/resources/jpg/5/2/1523885965125.jpg"></img>
                        <img width="28" height="28" style={{marginTop: '-40px'}} src="https://img.icons8.com/color/48/tiktok-verified-account.png" alt="tiktok-verified-account"/>
                    </div>

                    <h3 style={{color: '#000750', fontWeight: 'bold'}}>Daiane dos Santos</h3>

                    <div class="row text-center shadow" style={divCard}>
                        <div class="col-xs-12">

                            <div class="row"  style={{padding: '5px'}}>
                                <div class="col-xs-2" style={{paddingLeft: '15px'}}>
                                    <img style={icons} src="https://img.icons8.com/fluency/48/people-working-together.png"></img>
                                </div>
                                <div class="col-xs-10" style={{paddingTop: '10px', paddingLeft: '10px'}}>
                                    <span>Instituição</span>
                                </div>
                            </div>

                            <div class="row"  style={{padding: '5px'}}>
                                <div class="col-xs-2" style={{paddingLeft: '15px'}}>
                                    <img style={icons} src="https://img.icons8.com/fluency/48/add-phone.png"></img>
                                </div>
                                <div class="col-xs-10" style={{paddingTop: '10px', paddingLeft: '10px'}}>
                                    <span>(16) 99391-8691</span>
                                </div>
                            </div>

                            <div class="row"  style={{padding: '5px'}}>
                                <div class="col-xs-2" style={{paddingLeft: '15px'}}>
                                    <img style={icons} src="https://img.icons8.com/fluency/48/email.png"></img>
                                </div>
                                <div class="col-xs-10" style={{paddingTop: '10px', paddingLeft: '10px'}}>
                                    <span>exemplo@gmail.com</span>
                                </div>
                            </div>                                                     
                        </div>
                    </div>

                    <div class="row text-center"  style={{padding: '5px', paddingTop: '15px', display: 'flex', alignItems: 'center', justifyContent: 'center', marginBottom: '-10px'}}>
                        <div class="col-xs-2 text-center" style={{paddingLeft: '15px', }}>
                            <img style={{maxWidth: '40px'}} src="https://img.icons8.com/fluency/48/https://img.icons8.com/fluency/48/external-link.png"></img>
                        </div>
                        <div class="col-xs-10" style={{paddingTop: '10px', paddingLeft: '10px', paddingBottom: '1px', }}>
                            <h4>Acessar Perfil</h4>
                        </div>
                    </div>

                </div>
                <div class="card container col-8" style={{paddingLeft: '10%', backgroundColor: 'transparent', border: 'none', boxShadow: 'none'}}>
                    <h1 class="text-center" style={{fontColor: '#000750', fontWeight: 'bold'}}>Gostaria de conceder acesso?</h1>
                    <h5>Ao conceder acesso, será permitido que André visualize todos os seus documentos, sem restrições.</h5>

                    <div class="shadow" style={{marginTop: '25px', paddingTop: '20px', backgroundColor: '#000750', borderRadius: '10px', padding: '5px', color: 'white',}}>
                        <div class="row form-group" style={{padding: '25px', paddingLeft: '35px'}}>
                            <div>
                                <h4>Manter acesso por:</h4>
                                <input type="text" class="form-control" style={{maxWidth: '90%', marginTop: '20px'}} />
                                <div class="form-group">
                                <label for="exampleFormControlSelect1" style={{fontWeight: 'bold', marginTop: '10px'}}>Tempo</label>
                                    <select class="form-control" id="exampleFormControlSelect1" style={{maxWidth: '90%'}}>
                                        <option>Horas</option>
                                        <option>Dias</option>
                                        <option>Semanas</option>
                                        <option>Meses</option>
                                        <option>Anos</option>
                                    </select>
                                </div>                                
                            </div>

                            <div class="col" style={{fontSize: '20px', display: 'flex', alignItems: 'left', justifyContent: 'center', flexDirection: 'column', paddingLeft: '30px' }}>
                                <div class="form-check">
                                    <input class="form-check-input" type="checkbox" value="" id="defaultCheck1" />
                                    <label class="form-check-label" for="defaultCheck1">
                                        Aceito os termos de uso
                                    </label>
                                </div>
                                <div class="form-check">
                                    <input class="form-check-input" style={{ marginTop: '0' }} type="checkbox" value="" id="defaultCheck2" />
                                    <label class="form-check-label" for="defaultCheck2">
                                        Aceito a política de privacidade
                                    </label>
                                </div>
                                <div class="form-check">
                                    <input class="form-check-input" style={{ marginTop: '0' }} type="checkbox" value="" id="defaultCheck2" />
                                    <label class="form-check-label" for="defaultCheck2">
                                        Gostaria de receber e-mail quando o tempo expirar
                                    </label>
                                </div>
                            </div>
                        </div>                                               
                    </div>
                    <div style={{marginTop: '10px'}}>
                        <span>
                            Não nos responsábilizamos pelo uso indevido de seus documentos. Tenha cuidado ao compartilhar suas informações.
                        </span>
                    </div>
                    <div class="row btn-group">
                        <div class="btn text-center shadow" style={{marginLeft: '15px', marginTop: '25px', paddingTop: '20px', backgroundColor: '#000750', borderRadius: '10px', padding: '5px', color: 'white', width: '10%', }}>
                            <h4 style={{fontWeight: 'bold', justifyContent: 'center'}}>Permitir</h4>
                        </div>
                        <div class="btn text-center shadow" style={{marginLeft: '15px', marginTop: '25px', paddingTop: '20px', backgroundColor: '#00000', borderRadius: '10px', padding: '5px', color: '#000750', width: '7%', borderColor: '#000750', marginRight: '2%'}}>
                            <h4 style={{fontWeight: 'bold', }}>Negar</h4>
                        </div>
                    </div>
                </div>
            </div>
        </div>

    );

};

export default MyComponent