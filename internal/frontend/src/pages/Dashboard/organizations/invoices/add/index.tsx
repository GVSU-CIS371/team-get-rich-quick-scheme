import {DashboardLayout} from "@/components/ui/dashboard-layout";
import {useLocation, useRoute} from "preact-iso";
import {JSONForm} from "@/components/form/JSONForm";
import {JSONFormInput} from "@/components/form/JSONFormInput";
import {Button} from "@/components/ui/button";

export function AddInvoicePage() {
    const location = useLocation();
    const {params} = useRoute();

    const onSuccess = ({invoice}) => {
        location.route(`/dashboard/organizations/${params.id}/invoices/${invoice.ID}`);
    }

    return (
        <DashboardLayout>
            <div className="rounded-2xl border border-zinc-800 bg-zinc-950 p-6">
                <h2 className="text-lg font-medium">Add Invoice</h2>
                <JSONForm method="POST" action={`/api/v1/organizations/${params.id}/invoices`}
                          className="mt-4 space-y-3" onSuccess={onSuccess}>
                    <JSONFormInput name="note" label="Invoice Note" />
                    <Button>Submit</Button>
                </JSONForm>
            </div>
        </DashboardLayout>
    )
}