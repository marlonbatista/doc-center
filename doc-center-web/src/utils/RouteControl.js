import { Outlet, Navigate } from 'react-router-dom'

const PrivateRoutes = () => {
    let token = localStorage.getItem('token');
    return(
        token != null ? <Outlet/> : <Navigate to="/login"/>
    )
}

const PublicRoutes = () => {
    let token = localStorage.getItem('token');
    return(
        token == null ? <Outlet/> : <Navigate to="/"/>
    )
}

const RouteControl = {
    PrivateRoutes,
    PublicRoutes
}

export default RouteControl
