import { ref } from "vue";
export interface ExecutorOptions<T> {
    request: Promise<any>
    showError?: boolean
    isGlobalShowLoading?: boolean
    isDataObject?: boolean
    onResult?: (data: T) => void
    onError?: (error: any) => void
    onComplete?: () => void
  }


export function useExecutor() {
  const isLoading = ref(false)

  async function run<T>(options: ExecutorOptions<T>): Promise<void> {
    const {
      request,        
      onResult,         
      onError,           
      onComplete,
      showError, 
    } = options

    try {
      setLoading(true)
      const result = await request
      const data = await result.json()
      onResult?.(data)
    } catch (error) {
      const err = handleError(error, showError ?? true)
      onError?.(err)
    } finally {
      setLoading(false)
      onComplete?.()
    }
  }


  function setLoading(loading: boolean) {
    isLoading.value = loading;
    // const sharedStore = useSharedStore();
    // sharedStore.setGlobalLoading(loading, isGlobalShowLoading);
  }


  function handleError(error: any, showError: boolean) {
    alert(error.message);
    return {
      ...error,
      status: error.status,
      message: error.errors,
    };
  }

  return { run, isLoading };
}
