import {Dispatch, useContext, useEffect, useState} from "preact/hooks";
import axios, {AxiosInstance} from "axios";
import {createContext} from "preact";
import {StateUpdater} from "preact/hooks";

type Auth = {
    token: string;
    expiration: number;
}

const AuthContext = createContext(undefined);

interface AuthContextType {
    auth: Auth|null;
    setAuth: Dispatch<StateUpdater<Auth>>;
    authClient: AxiosInstance;
    setAuthClient: Dispatch<StateUpdater<AxiosInstance>>;
    isLoggedIn: () => boolean;
    saveToken: (token: string, expiration: number, apiOnly: boolean) => void;
    removeToken: () => void;
}

function loadAuth(): Auth|null {
    const token = localStorage.getItem("auth");
    const expiration = localStorage.getItem("auth-exp");

    if (!token || !expiration || isNaN(Number(expiration))) {
        return null;
    }

    if (Number(expiration) < Date.now()) {
        return null;
    }

    return {
        token,
        expiration: Number(expiration),
    }
}

export const AuthProvider = ({children}) => {
    const [auth, setAuth] = useState<Auth|null>(loadAuth());
    const [authClient, setAuthClient] = useState<AxiosInstance>(() => {
        const client = axios.create({
            validateStatus: (status) => {
                if (status === 401) {
                    removeToken();
                }
                return true;
            }
        });
        if (auth && auth.expiration > Date.now()) {
            client.defaults.headers['Authorization'] = auth.token;
        }
        return client;
    });

    useEffect(() => {
        if (auth && auth.token && auth.expiration > Date.now()) {
            authClient.defaults.headers['Authorization'] = auth.token;
        }
    }, [auth]);

    const isLoggedIn = (): boolean => {
        return auth && auth.expiration > Date.now();
    }

    const saveToken = (token: string, expiration: number):  void => {
        setAuth({
            token: token,
            expiration: expiration,
        });

        localStorage.setItem("auth", token);
        localStorage.setItem("auth-exp", String(expiration));
    }

    const removeToken = () => {
        localStorage.removeItem("auth");
        localStorage.removeItem("auth-exp");
        setAuth(null);
    }

    return (
        <AuthContext.Provider value={{
            auth, setAuth,
            authClient, setAuthClient,
            isLoggedIn, saveToken, removeToken
        }}>
            {children}
        </AuthContext.Provider>
    )
}

export const useAuth = (): AuthContextType => {
    const context = useContext(AuthContext);
    if (!context) {
        throw new Error('useAuth must be used within an AuthProvider');
    }
    return context;
}