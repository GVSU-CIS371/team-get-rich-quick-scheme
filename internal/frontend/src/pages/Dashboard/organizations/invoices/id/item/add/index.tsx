import {useLocation, useRoute} from "preact-iso";
import {DashboardLayout} from "@/components/ui/dashboard-layout";
import {JSONForm} from "@/components/form/JSONForm";
import {JSONFormInput} from "@/components/form/JSONFormInput";
import {Button} from "@/components/ui/button";

export function AddInvoiceItemPage() {
    const location = useLocation();
    const {params} = useRoute();
    const backRoute = `/dashboard/organizations/${params.id}/invoices/${params.invId}`;

    const onSuccess = () => {
        location.route(backRoute);
    }

    return (
        <DashboardLayout>
            <div className="rounded-2xl border border-zinc-800 bg-zinc-950 p-6">
                <h2 className="text-lg font-medium">Add Invoice Item</h2>
                <JSONForm method="POST" action={`/api/v1/organizations/${params.id}/invoices/${params.invId}/items`}
                          className="mt-4 space-y-3" onSuccess={onSuccess}>
                    <JSONFormInput name="description" label="Description" />
                    <JSONFormInput type="number" name="quantity" label="Quantity" />
                    <JSONFormInput type="number" name="price" label="Price" />
                    <Button>Submit</Button>
                    <Button onClick={() => location.route(backRoute)}>Go Back</Button>
                </JSONForm>
            </div>
        </DashboardLayout>
    )
}