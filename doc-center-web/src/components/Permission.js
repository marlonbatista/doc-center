import { borderRadius } from "@mui/system";
import React from "react";

const MyComponent = () => {
    const containerStyle = {
        maxWidth: '200px',
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
                <div class="col card text-center">

                    <span style={divCard}>Solicitar acesso</span>
                    <div class="row" style={{}}>
                        <img style={containerStyle} class="img img-fluid" src="https://img.estadao.com.br/resources/jpg/5/2/1523885965125.jpg"></img>
                    </div>

                    <h3 style={{color: '#000750'}}>Katia Exemplo</h3>

                    <div class="row text-center" style={divCard}>
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


                </div>
                <div class="card container col-8">
                    <h1>teste2</h1>
                </div>
            </div>
        </div>

    );

};

export default MyComponent