import {useAuth} from "@/components/auth/Auth";
import {useLocation} from "preact-iso";

export function LogoutPage() {
    const {removeToken} = useAuth();
    const location = useLocation();

    removeToken();
    location.route('/');

    return (
        <p>Logging out</p>
    )
}