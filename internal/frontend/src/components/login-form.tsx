import { cn } from "@/lib/utils"
import { Button } from "@/components/ui/button"
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"
import {
  Field,
  FieldDescription,
  FieldGroup,
} from "@/components/ui/field"
import {JSONForm} from "@/components/form/JSONForm";
import {JSONFormInput} from "@/components/form/JSONFormInput";
import {useAuth} from "@/components/auth/Auth";
import {useState} from "preact/hooks";
import {useLocation} from "preact-iso";
import {AlertCircleIcon} from "lucide-react";
import {Alert, AlertDescription, AlertTitle} from "@/components/ui/alert";

export function LoginForm({
  className,
  ...props
}: React.ComponentProps<"div">) {
  const {isLoggedIn, saveToken} = useAuth();
  const [error, setError] = useState<string>('');
  const [errorVisible, setErrorVisible] = useState<boolean>(false);
  const location = useLocation();

  if (isLoggedIn()) {
    location.route('/dashboard');
    return;
  }

  const onSuccess = ({session, expiration}) => {
    saveToken(session, expiration);
    location.route('/dashboard');
  }

  const onError = (errors: string[]) => {
    setError(errors[0]);
    setErrorVisible(true);
  }

  return (
    <div className={cn("flex flex-col gap-6", className)} {...props}>
      <Card>
        <CardHeader>
          <CardTitle>Login to your account</CardTitle>
          <CardDescription>
            Enter your email below to login to your account
          </CardDescription>
        </CardHeader>
        <CardContent>
          <JSONForm method="POST" action="/api/v1/login" onFail={onError} onSuccess={onSuccess}>
            {errorVisible && <Alert variant="destructive" className="my-3">
              <AlertCircleIcon />
              <AlertTitle>Error</AlertTitle>
              <AlertDescription>
                {error}
              </AlertDescription>
            </Alert>}
            <FieldGroup>
              <JSONFormInput
                  type="email"
                  name="email"
                  label="Email"
                  required
              />
              <JSONFormInput
                  type="password"
                  name="password"
                  label="Password"
                  required
              />
              <Field>
                <Button type="submit">Login</Button>
                <FieldDescription className="text-center">
                  Don&apos;t have an account? <a href="/register">Sign up</a>
                </FieldDescription>
              </Field>
            </FieldGroup>
          </JSONForm>
        </CardContent>
      </Card>
    </div>
  )
}
