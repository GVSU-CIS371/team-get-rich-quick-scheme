import {DashboardLayout} from "@/components/ui/dashboard-layout";
import {JSONForm} from "@/components/form/JSONForm";
import {JSONFormInput} from "@/components/form/JSONFormInput";
import {Button} from "@/components/ui/button";
import {useState} from "preact/hooks";

export function SettingsPage() {
    const [changed, setChanged] = useState(false);

    return (
        <DashboardLayout>
            <div className="rounded-2xl border border-zinc-800 bg-zinc-950 p-6">
                <h2 className="text-lg font-medium">Change Password</h2>

                {changed && (
                    <p>Your password has been updated!</p>
                )}

                <JSONForm method="PUT" action="/api/v1/user/password" className="mt-4 space-y-3"
                          onSuccess={() => setChanged(true)}>
                    <JSONFormInput type="password" name="currentPassword" label="Current Password" />
                    <JSONFormInput type="password" name="newPassword" label="New Password" />
                    <Button>Submit</Button>
                </JSONForm>
            </div>
        </DashboardLayout>
    )
}