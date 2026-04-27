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
import {Alert, AlertDescription, AlertTitle} from "@/components/ui/alert";
import {AlertCircleIcon} from "lucide-react";

export function RegisterForm({
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
          <CardTitle>Register for an account</CardTitle>
          <CardDescription>
            Register now to join our platform
          </CardDescription>
        </CardHeader>
        <CardContent>
          <JSONForm method="POST" action="/api/v1/register" onFail={onError} onSuccess={onSuccess}>
            {errorVisible && <Alert variant="destructive" className="my-3">
              <AlertCircleIcon />
              <AlertTitle>Error</AlertTitle>
              <AlertDescription>
                {error}
              </AlertDescription>
            </Alert>}

            <FieldGroup>
              <JSONFormInput
                  type="text"
                  name="firstName"
                  label="First Name"
                  required
              />
              <JSONFormInput
                  type="text"
                  name="lastName"
                  label="Last Name"
                  required
              />
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
                <Button type="submit">Register</Button>
                <FieldDescription className="text-center">
                  Already have an account? <a href="/login">Login</a>
                </FieldDescription>
              </Field>
            </FieldGroup>
          </JSONForm>
        </CardContent>
      </Card>
    </div>
  )
}
