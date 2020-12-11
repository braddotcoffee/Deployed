import { Deployment } from '../deployment_pb';
import DeployStatus = Deployment.DeployStatus;

export function readableStatus(status: number): string {
    if (status === DeployStatus.IN_PROGRESS) {
        return 'In Progress';
    }
    if (status === DeployStatus.NOT_STARTED) {
        return 'Not Started';
    }
    if (status === DeployStatus.ERROR) {
        return 'Error';
    }
    if (status === DeployStatus.COMPLETE) {
        return 'Complete';
    }
    return 'Unknown Status';
}
