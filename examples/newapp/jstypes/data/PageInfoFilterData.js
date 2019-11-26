
	import {PageInfoFilter, PageInfoFilterFilter} from '../apiModel';

let pageInfoFilterData = function () {
    return {
        pageInfoFilterFilter: new PageInfoFilterFilter(),
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
        currentPageInfoFilterItem: {
            item: new PageInfoFilter(),
            hasChange: false,
            isSelected: false,
            canDelete: false,
            showDeleteConfirmation: false,
        },
    }
}();

export default pageInfoFilterData;

