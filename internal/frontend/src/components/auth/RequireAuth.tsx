import {useAuth} from "./Auth";
import {useLocation} from "preact-iso";

export function RequireAuth({children}) {
    const {isLoggedIn} = useAuth();
    const location = useLocation();

    if (!isLoggedIn()) {
        location.route('/login');
    }

    return (
        <>
            {children}
        </>
    )
}