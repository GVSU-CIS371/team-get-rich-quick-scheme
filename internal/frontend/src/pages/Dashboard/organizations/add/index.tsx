import {DashboardLayout} from "@/components/ui/dashboard-layout";
import {JSONForm} from "@/components/form/JSONForm";
import {JSONFormInput} from "@/components/form/JSONFormInput";
import {Button} from "@/components/ui/button";
import {useLocation} from "preact-iso";

export function AddOrganizationPage() {
    const location = useLocation();

    const onSuccess = ({organization}) => {
        location.route(`/dashboard/organizations/${organization.ID}`);
    }

    return (
        <DashboardLayout>
            <div className="rounded-2xl border border-zinc-800 bg-zinc-950 p-6">
                <h2 className="text-lg font-medium">Add Organization</h2>
                <JSONForm method="POST" action="/api/v1/organizations" className="mt-4 space-y-3" onSuccess={onSuccess}>
                    <JSONFormInput name="name" label="Organization Name" />
                    <JSONFormInput name="description" label="Description" />
                    <Button>Submit</Button>
                </JSONForm>
            </div>
        </DashboardLayout>
    )
}