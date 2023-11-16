import { toast } from 'react-toastify';

export function alert(type, message) {
  switch (type) {
    case 'success':
      toast.success(message, { position: toast.POSITION.TOP_CENTER });
      break;
    case 'error':
      toast.error(message, { position: toast.POSITION.TOP_CENTER });
      break;
    case 'warn':
      toast.warn(message, { position: toast.POSITION.TOP_CENTER });
      break;
      case 'info':
        toast.info(message, { position: toast.POSITION.TOP_CENTER });
        break;
    default:
      toast(message, { position: toast.POSITION.TOP_CENTER });
      break;
  }
}