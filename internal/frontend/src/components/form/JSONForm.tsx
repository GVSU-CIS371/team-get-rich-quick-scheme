import {createContext, JSX, TargetedEvent} from "preact";
import {useAuth} from "../auth/Auth";
import {useContext, useState} from "preact/hooks";

const JSONFormContext = createContext(undefined);

interface JSONFormContextType {
    issues: Record<string, string>
}

type JSONFormProps = JSX.IntrinsicElements["form"] & {
    method: string;
    action: string;
    onSuccess?: (data: any) => void;
    onFail?: (errors: string[]) => void;
    onIssues?: (issues: Record<string, string>) => void;
};

export function JSONForm({children, method, action, onSuccess, onFail, onIssues, ...props}: JSONFormProps) {
    const {authClient} = useAuth();
    const [issues, setIssues] = useState<Record<string, string>>({});

    const onSubmit = async (e: TargetedEvent<HTMLFormElement, Event>) => {
        e.preventDefault();

        const data = new FormData(e.currentTarget);

        let object = {};
        data.forEach((value, key) => object[key] = value);
        const json = JSON.stringify(object);

        const res = await authClient.post(action, JSON.parse(json));

        if (res.status === 200) {
            onSuccess?.(res.data.data);
            return;
        }

        if (res.data.issues) {
            setIssues(res.data.issues);
            onIssues?.(res.data.issues);
            return;
        } else {
            setIssues({});
        }

        if (res.data.errors) {
            onFail?.(res.data.errors);
            return;
        }
    }

    return (
        <JSONFormContext.Provider value={{
            issues
        }}>
            <form onSubmit={onSubmit} {...props}>
                {children}
            </form>
        </JSONFormContext.Provider>
    )
}

export const useJSONForm = (): JSONFormContextType => {
    const context = useContext(JSONFormContext);
    if (!context) {
        throw new Error('useJSONForm must be used within an JSONForm');
    }
    return context;
}