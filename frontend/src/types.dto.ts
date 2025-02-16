export interface IListItem {
    id: number
    title: string
    is_done: boolean
  }
  
export interface IState {
    list: IListItem[]
    newTask: string
}
  