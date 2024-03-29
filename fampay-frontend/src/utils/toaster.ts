import { Id, ToastOptions, UpdateOptions, toast } from 'react-toastify';

class Toaster {
  static toastSettings: ToastOptions = {
    position: 'top-center',
    autoClose: 1000,
    hideProgressBar: false,
    closeOnClick: true,
    pauseOnHover: false,
    draggable: true,
    progress: undefined,
    theme: 'light',
  };

  static success(message: string, toastId: string = 'default'): Id {
    return toast.success(message, { ...this.toastSettings, toastId });
  }

  static error(message: string, toastId: string = 'default'): Id {
    return toast.error(message, { ...this.toastSettings, toastId });
  }

  static startLoad(message: string = 'Loading...', toastId: string = 'default'): Id {
    return toast.loading(message, { ...this.toastSettings, toastId });
  }

  static stopLoad(loader: Id, message: string, res: number): void {
    const settings: UpdateOptions = { ...this.toastSettings };
    settings.render = message;
    settings.type = res === 1 ? 'success' : 'error';
    settings.isLoading = false;
    toast.update(loader, settings);
  }
}

export default Toaster;
