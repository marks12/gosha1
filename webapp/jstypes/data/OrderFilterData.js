
	import {OrderFilter, OrderFilterFilter} from '../apiModel';

let orderFilterData = function () {
    return {
        orderFilterFilter: new OrderFilterFilter(),
        panel: {
            type: '',
            show: false,
            create: 'create',
            edit: 'edit'
        },
        panelHeaderCreate: 'Создать',
        panelHeaderEdit: 'Изменить',
        panelSubmitButtonTextCreate: 'Создать',
        panelSubmitButtonTextEdit: 'Изменить',
        currentOrderFilterItem: {
            item: new OrderFilter(),
            hasChange: false,
            isSelected: false,
            canDelete: false,
            showDeleteConfirmation: false,
        },
    }
}();

export default orderFilterData;

