import {useJSONForm} from "@/components/form/JSONForm";
import {Field, FieldDescription, FieldLabel} from "@/components/ui/field";
import { Input } from "@/components/ui/input";

export function JSONFormInput({name, label, hidden, ...props}: {name: string, label: string, [props: string]: any}) {
    const {issues} = useJSONForm();

    return (
        <Field data-invalid={name in issues}>
            <FieldLabel htmlFor={name}>{label}</FieldLabel>
            <Input id={name} name={name} aria-invalid={name in issues} aria-label={label} {...props} />
            {!hidden && name in issues && <FieldDescription>{issues[name]}</FieldDescription>}
        </Field>
    )
}