import {RegisterForm} from "@/components/register-form";

export function Register() {
    return (
        <div class="absolute w-full top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2">
            <div className="flex w-full justify-center">
                <RegisterForm className="w-full mx-3 md:mx-0 md:w-140" />
            </div>
        </div>
    )
}