#include "cdrv.h"
#include <QtCore/QObject>
#include "qtapp.h"
#include <QWidget>
#include <QMenuBar>
#include <QMenu>
#include <QAction>
#include <QTabWidget>
#include <QLayout>
#include <QBoxLayout>
#include <QHBoxLayout>
#include <QVBoxLayout>
#include <QAbstractButton>
#include <QPushButton>
#include <QCheckBox>
#include <QRadioButton>
#include <QFrame>
#include <QLabel>
#include <QGroupBox>
#include <QDialog>
#include <QComboBox>
#include <QPainter>
// drvclass enums
enum DRVCLASS_ENUMS {
    DRVCLASS_NONE = 0,
    DRVCLASS_OBJECT,
    DRVCLASS_APP,
    DRVCLASS_WIDGET,
    DRVCLASS_MENUBAR,
    DRVCLASS_MENU,
    DRVCLASS_ACTION,
    DRVCLASS_TABWIDGET,
    DRVCLASS_LAYOUT,
    DRVCLASS_BOXLAYOUT,
    DRVCLASS_HBOXLAYOUT,
    DRVCLASS_VBOXLAYOUT,
    DRVCLASS_BASEBUTTON,
    DRVCLASS_BUTTON,
    DRVCLASS_CHECKBOX,
    DRVCLASS_RADIO,
    DRVCLASS_FRAME,
    DRVCLASS_LABEL,
    DRVCLASS_GROUPBOX,
    DRVCLASS_DIALOG,
    DRVCLASS_COMBOBOX,
    DRVCLASS_PAINTER,
    DRVCLASS_LAST
};
// DRVCLASS_OBJECT drvid enums
enum OBJECT_ENUMS {
    OBJECT_NONE = 0,
    OBJECT_DESTROY,
    OBJECT_STARTTIMER,
    OBJECT_KILLTIMER,
    OBJECT_ONTIMEREVENT,
    OBJECT_LAST
};
// DRVCLASS_APP drvid enums
enum APP_ENUMS {
    APP_NONE = 0,
    APP_INIT,
    APP_CLOSE,
    APP_EXEC,
    APP_EXIT,
    APP_SETFONT,
    APP_FONT,
    APP_CLOSEALLWINDOWS,
    APP_ONREMOVEITEM,
    APP_LAST
};
// DRVCLASS_WIDGET drvid enums
enum WIDGET_ENUMS {
    WIDGET_NONE = 0,
    WIDGET_INIT,
    WIDGET_DESTROY,
    WIDGET_SETAUTODESTROY,
    WIDGET_AUTODESTROY,
    WIDGET_SETLAYOUT,
    WIDGET_LAYOUT,
    WIDGET_SETPARENT,
    WIDGET_PARENT,
    WIDGET_SETVISIBLE,
    WIDGET_ISVISIBLE,
    WIDGET_SETWINDOWTITLE,
    WIDGET_WINDOWTITLE,
    WIDGET_SETPOS,
    WIDGET_POS,
    WIDGET_SETSIZE,
    WIDGET_SIZE,
    WIDGET_SETGEOMETRY,
    WIDGET_GEOMETRY,
    WIDGET_SETFONT,
    WIDGET_FONT,
    WIDGET_CLOSE,
    WIDGET_UPDATE,
    WIDGET_REPAINT,
    WIDGET_ONSHOWEVENT,
    WIDGET_ONHIDEEVENT,
    WIDGET_ONCLOSEEVENT,
    WIDGET_ONKEYPRESSEVENT,
    WIDGET_ONKEYRELEASEEVENT,
    WIDGET_ONMOUSEPRESSEVENT,
    WIDGET_ONMOUSERELEASEEVENT,
    WIDGET_ONMOUSEMOVEEVENT,
    WIDGET_ONMOUSEDOUBLECLICKEVENT,
    WIDGET_ONMOVEEVENT,
    WIDGET_ONPAINTEVENT,
    WIDGET_ONRESIZEEVENT,
    WIDGET_ONENTEREVENT,
    WIDGET_ONLEAVEEVENT,
    WIDGET_ONFOCUSINEVENT,
    WIDGET_ONFOCUSOUTEVENT,
    WIDGET_LAST
};
// DRVCLASS_MENUBAR drvid enums
enum MENUBAR_ENUMS {
    MENUBAR_NONE = 0,
    MENUBAR_INIT,
    MENUBAR_ADDMENU,
    MENUBAR_LAST
};
// DRVCLASS_MENU drvid enums
enum MENU_ENUMS {
    MENU_NONE = 0,
    MENU_INIT,
    MENU_SETTITLE,
    MENU_TITLE,
    MENU_ADDACTION,
    MENU_LAST
};
// DRVCLASS_ACTION drvid enums
enum ACTION_ENUMS {
    ACTION_NONE = 0,
    ACTION_INIT,
    ACTION_SETTEXT,
    ACTION_TEXT,
    ACTION_ONTRIGGERED,
    ACTION_LAST
};
// DRVCLASS_TABWIDGET drvid enums
enum TABWIDGET_ENUMS {
    TABWIDGET_NONE = 0,
    TABWIDGET_INIT,
    TABWIDGET_ADDTAB,
    TABWIDGET_CLEAR,
    TABWIDGET_COUNT,
    TABWIDGET_CURRENTINDEX,
    TABWIDGET_CURRENTWIDGET,
    TABWIDGET_SETCURRENTINDEX,
    TABWIDGET_SETCURRENTWIDGET,
    TABWIDGET_INDEXOF,
    TABWIDGET_INSERTTAB,
    TABWIDGET_REMOVETAB,
    TABWIDGET_SETTABTEXT,
    TABWIDGET_SETTABTOOLTIP,
    TABWIDGET_TABTEXT,
    TABWIDGET_TABTOOLTIP,
    TABWIDGET_WIDGETOF,
    TABWIDGET_ONCURRENTCHANGED,
    TABWIDGET_LAST
};
// DRVCLASS_LAYOUT drvid enums
enum LAYOUT_ENUMS {
    LAYOUT_NONE = 0,
    LAYOUT_ADDLAYOUT,
    LAYOUT_ADDWIDGET,
    LAYOUT_SETSPACING,
    LAYOUT_SPACING,
    LAYOUT_SETMARGINS,
    LAYOUT_MARGINS,
    LAYOUT_SETMENUBAR,
    LAYOUT_MENUBAR,
    LAYOUT_LAST
};
// DRVCLASS_BOXLAYOUT drvid enums
enum BOXLAYOUT_ENUMS {
    BOXLAYOUT_NONE = 0,
    BOXLAYOUT_ADDLAYOUTWITH,
    BOXLAYOUT_ADDWIDGETWITH,
    BOXLAYOUT_ADDSPACING,
    BOXLAYOUT_ADDSTRETCH,
    BOXLAYOUT_LAST
};
// DRVCLASS_HBOXLAYOUT drvid enums
enum HBOXLAYOUT_ENUMS {
    HBOXLAYOUT_NONE = 0,
    HBOXLAYOUT_INIT,
    HBOXLAYOUT_LAST
};
// DRVCLASS_VBOXLAYOUT drvid enums
enum VBOXLAYOUT_ENUMS {
    VBOXLAYOUT_NONE = 0,
    VBOXLAYOUT_INIT,
    VBOXLAYOUT_LAST
};
// DRVCLASS_BASEBUTTON drvid enums
enum BASEBUTTON_ENUMS {
    BASEBUTTON_NONE = 0,
    BASEBUTTON_SETTEXT,
    BASEBUTTON_TEXT,
    BASEBUTTON_LAST
};
// DRVCLASS_BUTTON drvid enums
enum BUTTON_ENUMS {
    BUTTON_NONE = 0,
    BUTTON_INIT,
    BUTTON_SETFLAT,
    BUTTON_ISFLAT,
    BUTTON_SETDEFAULT,
    BUTTON_ISDEFAULT,
    BUTTON_SETMENU,
    BUTTON_MENU,
    BUTTON_ONCLICKED,
    BUTTON_LAST
};
// DRVCLASS_CHECKBOX drvid enums
enum CHECKBOX_ENUMS {
    CHECKBOX_NONE = 0,
    CHECKBOX_INIT,
    CHECKBOX_SETCHECK,
    CHECKBOX_CHECK,
    CHECKBOX_SETTRISTATE,
    CHECKBOX_ISTRISTATE,
    CHECKBOX_ONSTATECHANGED,
    CHECKBOX_LAST
};
// DRVCLASS_RADIO drvid enums
enum RADIO_ENUMS {
    RADIO_NONE = 0,
    RADIO_INIT,
    RADIO_ONCLICKED,
    RADIO_LAST
};
// DRVCLASS_FRAME drvid enums
enum FRAME_ENUMS {
    FRAME_NONE = 0,
    FRAME_INIT,
    FRAME_SETFRAMESTYLE,
    FRAME_FRAMESTYLE,
    FRAME_SETFRAMERECT,
    FRAME_FRAMERECT,
    FRAME_LAST
};
// DRVCLASS_LABEL drvid enums
enum LABEL_ENUMS {
    LABEL_NONE = 0,
    LABEL_INIT,
    LABEL_SETTEXT,
    LABEL_TEXT,
    LABEL_SETWORDWRAP,
    LABEL_WORDWRAP,
    LABEL_SETTEXTFORMAT,
    LABEL_TEXTFORMAT,
    LABEL_ONLINKACTIVATED,
    LABEL_LAST
};
// DRVCLASS_GROUPBOX drvid enums
enum GROUPBOX_ENUMS {
    GROUPBOX_NONE = 0,
    GROUPBOX_INIT,
    GROUPBOX_SETTITLE,
    GROUPBOX_TITLE,
    GROUPBOX_LAST
};
// DRVCLASS_DIALOG drvid enums
enum DIALOG_ENUMS {
    DIALOG_NONE = 0,
    DIALOG_INIT,
    DIALOG_SETMODAL,
    DIALOG_ISMODAL,
    DIALOG_SETRESULT,
    DIALOG_RESULT,
    DIALOG_EXEC,
    DIALOG_DONE,
    DIALOG_ACCEPT,
    DIALOG_REJECT,
    DIALOG_ONACCEPTED,
    DIALOG_ONREJECTED,
    DIALOG_LAST
};
// DRVCLASS_COMBOBOX drvid enums
enum COMBOBOX_ENUMS {
    COMBOBOX_NONE = 0,
    COMBOBOX_INIT,
    COMBOBOX_COUNT,
    COMBOBOX_SETCURRENTINDEX,
    COMBOBOX_CURRENTINDEX,
    COMBOBOX_CURRENTTEXT,
    COMBOBOX_SETEDITABLE,
    COMBOBOX_ISEDITABLE,
    COMBOBOX_SETMAXCOUNT,
    COMBOBOX_MAXCOUNT,
    COMBOBOX_SETMAXVISIBLEITEMS,
    COMBOBOX_MAXVISIBLEITEMS,
    COMBOBOX_SETMINIMUMCONTENTSLENGTH,
    COMBOBOX_MINIMUNCONTENTSLENGHT,
    COMBOBOX_ADDITEM,
    COMBOBOX_INSERTITEM,
    COMBOBOX_REMOVEITEM,
    COMBOBOX_ITEMTEXT,
    COMBOBOX_ONCURRENTINDEXCHANGED,
    COMBOBOX_LAST
};
// DRVCLASS_PAINTER drvid enums
enum PAINTER_ENUMS {
    PAINTER_NONE = 0,
    PAINTER_INIT,
    PAINTER_DESTROY,
    PAINTER_BEGIN,
    PAINTER_END,
    PAINTER_SETFONT,
    PAINTER_FONT,
    PAINTER_DRAWPOINT,
    PAINTER_DRAWLINE,
    PAINTER_DRAWLINES,
    PAINTER_DRAWPOLYLINE,
    PAINTER_DRAWTEXT,
    PAINTER_LAST
};
int drv_object(int drvid, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6, void* a7, void* a8, void* a9)
{
    handle_head* head = (handle_head*)a0;
    QObject *self = (QObject*)head->native;
    switch (drvid) {
    case OBJECT_DESTROY: {
        drvDelObj(a0,self);
        break;
    }
    case OBJECT_STARTTIMER: {
        drvSetInt(a2,self->startTimer(drvGetInt(a1)));
        break;
    }
    case OBJECT_KILLTIMER: {
        self->killTimer(drvGetInt(a1));
        break;
    }
    case OBJECT_ONTIMEREVENT: {
        drvNewEvent(QEvent::Timer,a0,a1,a2);
        break;
    }
    default:
        return 0;
    }
    return 1;
}

int drv_app(int drvid, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6, void* a7, void* a8, void* a9)
{
    handle_head* head = (handle_head*)a0;
    QtApp *self = (QtApp*)head->native;
    switch (drvid) {
    case APP_INIT: {
        drvNewObj(a0,new QtApp);
        break;
    }
    case APP_CLOSE: {
        drvDelObj(a0,self);
        break;
    }
    case APP_EXEC: {
        drvSetInt(a1,self->exec());
        break;
    }
    case APP_EXIT: {
        self->exit(drvGetInt(a1));
        break;
    }
    case APP_SETFONT: {
        self->setFont(drvGetFont(a1));
        break;
    }
    case APP_FONT: {
        drvSetFont(a1,self->font());
        break;
    }
    case APP_CLOSEALLWINDOWS: {
        self->closeAllWindows();
        break;
    }
    case APP_ONREMOVEITEM: {
        self->pfnDeleteObj = a1;
        break;
    }
    default:
        return 0;
    }
    return 1;
}

int drv_widget(int drvid, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6, void* a7, void* a8, void* a9)
{
    handle_head* head = (handle_head*)a0;
    QWidget *self = (QWidget*)head->native;
    switch (drvid) {
    case WIDGET_INIT: {
        drvNewObj(a0,new QWidget);
        break;
    }
    case WIDGET_DESTROY: {
        drvDelObj(a0,self);
        break;
    }
    case WIDGET_SETAUTODESTROY: {
        self->setAttribute(Qt::WA_DeleteOnClose,drvGetBool(a1));
        break;
    }
    case WIDGET_AUTODESTROY: {
        drvSetBool(a1,self->testAttribute(Qt::WA_DeleteOnClose));
        break;
    }
    case WIDGET_SETLAYOUT: {
        self->setLayout(drvGetLayout(a1));
        break;
    }
    case WIDGET_LAYOUT: {
        drvSetHandle(a1,self->layout());
        break;
    }
    case WIDGET_SETPARENT: {
        self->setParent(drvGetWidget(a1));
        break;
    }
    case WIDGET_PARENT: {
        drvSetHandle(a1,self->parent());
        break;
    }
    case WIDGET_SETVISIBLE: {
        self->setVisible(drvGetBool(a1));
        break;
    }
    case WIDGET_ISVISIBLE: {
        drvSetBool(a1,self->isVisible());
        break;
    }
    case WIDGET_SETWINDOWTITLE: {
        self->setWindowTitle(drvGetString(a1));
        break;
    }
    case WIDGET_WINDOWTITLE: {
        drvSetString(a1,self->windowTitle());
        break;
    }
    case WIDGET_SETPOS: {
        self->move(drvGetPoint(a1));
        break;
    }
    case WIDGET_POS: {
        drvSetPoint(a1,self->pos());
        break;
    }
    case WIDGET_SETSIZE: {
        self->resize(drvGetSize(a1));
        break;
    }
    case WIDGET_SIZE: {
        drvSetSize(a1,self->size());
        break;
    }
    case WIDGET_SETGEOMETRY: {
        self->setGeometry(drvGetRect(a1));
        break;
    }
    case WIDGET_GEOMETRY: {
        drvSetRect(a1,self->geometry());
        break;
    }
    case WIDGET_SETFONT: {
        self->setFont(drvGetFont(a1));
        break;
    }
    case WIDGET_FONT: {
        drvSetFont(a1,self->font());
        break;
    }
    case WIDGET_CLOSE: {
        self->close();
        break;
    }
    case WIDGET_UPDATE: {
        self->update();
        break;
    }
    case WIDGET_REPAINT: {
        self->repaint();
        break;
    }
    case WIDGET_ONSHOWEVENT: {
        drvNewEvent(QEvent::Show,a0,a1,a2);
        break;
    }
    case WIDGET_ONHIDEEVENT: {
        drvNewEvent(QEvent::Hide,a0,a1,a2);
        break;
    }
    case WIDGET_ONCLOSEEVENT: {
        drvNewEvent(QEvent::Close,a0,a1,a2);
        break;
    }
    case WIDGET_ONKEYPRESSEVENT: {
        drvNewEvent(QEvent::KeyPress,a0,a1,a2);
        break;
    }
    case WIDGET_ONKEYRELEASEEVENT: {
        drvNewEvent(QEvent::KeyRelease,a0,a1,a2);
        break;
    }
    case WIDGET_ONMOUSEPRESSEVENT: {
        drvNewEvent(QEvent::MouseButtonPress,a0,a1,a2);
        break;
    }
    case WIDGET_ONMOUSERELEASEEVENT: {
        drvNewEvent(QEvent::MouseButtonRelease,a0,a1,a2);
        break;
    }
    case WIDGET_ONMOUSEMOVEEVENT: {
        drvNewEvent(QEvent::MouseMove,a0,a1,a2);
        break;
    }
    case WIDGET_ONMOUSEDOUBLECLICKEVENT: {
        drvNewEvent(QEvent::MouseButtonDblClick,a0,a1,a2);
        break;
    }
    case WIDGET_ONMOVEEVENT: {
        drvNewEvent(QEvent::Move,a0,a1,a2);
        break;
    }
    case WIDGET_ONPAINTEVENT: {
        drvNewEvent(QEvent::Paint,a0,a1,a2);
        break;
    }
    case WIDGET_ONRESIZEEVENT: {
        drvNewEvent(QEvent::Resize,a0,a1,a2);
        break;
    }
    case WIDGET_ONENTEREVENT: {
        drvNewEvent(QEvent::Enter,a0,a1,a2);
        break;
    }
    case WIDGET_ONLEAVEEVENT: {
        drvNewEvent(QEvent::Leave,a0,a1,a2);
        break;
    }
    case WIDGET_ONFOCUSINEVENT: {
        drvNewEvent(QEvent::FocusIn,a0,a1,a2);
        break;
    }
    case WIDGET_ONFOCUSOUTEVENT: {
        drvNewEvent(QEvent::FocusOut,a0,a1,a2);
        break;
    }
    default:
        return 0;
    }
    return 1;
}

int drv_menubar(int drvid, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6, void* a7, void* a8, void* a9)
{
    handle_head* head = (handle_head*)a0;
    QMenuBar *self = (QMenuBar*)head->native;
    switch (drvid) {
    case MENUBAR_INIT: {
        drvNewObj(a0,new QMenuBar);
        break;
    }
    case MENUBAR_ADDMENU: {
        self->addMenu(drvGetMenu(a1));
        break;
    }
    default:
        return 0;
    }
    return 1;
}

int drv_menu(int drvid, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6, void* a7, void* a8, void* a9)
{
    handle_head* head = (handle_head*)a0;
    QMenu *self = (QMenu*)head->native;
    switch (drvid) {
    case MENU_INIT: {
        drvNewObj(a0,new QMenu);
        break;
    }
    case MENU_SETTITLE: {
        self->setTitle(drvGetString(a1));
        break;
    }
    case MENU_TITLE: {
        drvSetString(a1,self->title());
        break;
    }
    case MENU_ADDACTION: {
        self->addAction(drvGetAction(a1));
        break;
    }
    default:
        return 0;
    }
    return 1;
}

int drv_action(int drvid, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6, void* a7, void* a8, void* a9)
{
    handle_head* head = (handle_head*)a0;
    QAction *self = (QAction*)head->native;
    switch (drvid) {
    case ACTION_INIT: {
        drvNewObj(a0, new QAction(qApp));
        break;
    }
    case ACTION_SETTEXT: {
        self->setText(drvGetString(a1));
        break;
    }
    case ACTION_TEXT: {
        drvSetString(a1,self->text());
        break;
    }
    case ACTION_ONTRIGGERED: {
        QObject::connect(self,SIGNAL(triggered(bool)),drvNewSignal(self,a1,a2),SLOT(call(bool)));
        break;
    }
    default:
        return 0;
    }
    return 1;
}

int drv_tabwidget(int drvid, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6, void* a7, void* a8, void* a9)
{
    handle_head* head = (handle_head*)a0;
    QTabWidget *self = (QTabWidget*)head->native;
    switch (drvid) {
    case TABWIDGET_INIT: {
        drvNewObj(a0,new QTabWidget);
        break;
    }
    case TABWIDGET_ADDTAB: {
        self->addTab(drvGetWidget(a1),drvGetString(a2));
        break;
    }
    case TABWIDGET_CLEAR: {
        self->clear();
        break;
    }
    case TABWIDGET_COUNT: {
        drvSetInt(a1,self->count());
        break;
    }
    case TABWIDGET_CURRENTINDEX: {
        drvSetInt(a1,self->currentIndex());
        break;
    }
    case TABWIDGET_CURRENTWIDGET: {
        drvSetHandle(a1,self->currentWidget());
        break;
    }
    case TABWIDGET_SETCURRENTINDEX: {
        self->setCurrentIndex(drvGetInt(a1));
        break;
    }
    case TABWIDGET_SETCURRENTWIDGET: {
        self->setCurrentWidget(drvGetWidget(a1));
        break;
    }
    case TABWIDGET_INDEXOF: {
        drvSetInt(a2,self->indexOf(drvGetWidget(a1)));
        break;
    }
    case TABWIDGET_INSERTTAB: {
        self->insertTab(drvGetInt(a1),drvGetWidget(a2),drvGetString(a3));
        break;
    }
    case TABWIDGET_REMOVETAB: {
        self->removeTab(drvGetInt(a1));
        break;
    }
    case TABWIDGET_SETTABTEXT: {
        self->setTabText(drvGetInt(a1),drvGetString(a2));
        break;
    }
    case TABWIDGET_SETTABTOOLTIP: {
        self->setTabToolTip(drvGetInt(a1),drvGetString(a2));
        break;
    }
    case TABWIDGET_TABTEXT: {
        drvSetString(a2,self->tabText(drvGetInt(a1)));
        break;
    }
    case TABWIDGET_TABTOOLTIP: {
        drvSetString(a2,self->tabToolTip(drvGetInt(a1)));
        break;
    }
    case TABWIDGET_WIDGETOF: {
        drvSetHandle(a2,self->widget(drvGetInt(a1)));
        break;
    }
    case TABWIDGET_ONCURRENTCHANGED: {
        QObject::connect(self,SIGNAL(currentChanged(int)),drvNewSignal(self,a1,a2),SLOT(call(int)));
        break;
    }
    default:
        return 0;
    }
    return 1;
}

int drv_layout(int drvid, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6, void* a7, void* a8, void* a9)
{
    handle_head* head = (handle_head*)a0;
    QLayout *self = (QLayout*)head->native;
    switch (drvid) {
    case LAYOUT_ADDLAYOUT: {
        self->addItem(drvGetLayout(a1));
        break;
    }
    case LAYOUT_ADDWIDGET: {
        self->addWidget(drvGetWidget(a1));
        break;
    }
    case LAYOUT_SETSPACING: {
        self->setSpacing(drvGetInt(a1));
        break;
    }
    case LAYOUT_SPACING: {
        drvSetInt(a1,self->spacing());
        break;
    }
    case LAYOUT_SETMARGINS: {
        self->setContentsMargins(drvGetMargins(a1));
        break;
    }
    case LAYOUT_MARGINS: {
        drvSetMargins(a1,self->contentsMargins());
        break;
    }
    case LAYOUT_SETMENUBAR: {
        self->setMenuBar(drvGetWidget(a1));
        break;
    }
    case LAYOUT_MENUBAR: {
        drvSetHandle(a1,self->menuBar());
        break;
    }
    default:
        return 0;
    }
    return 1;
}

int drv_boxlayout(int drvid, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6, void* a7, void* a8, void* a9)
{
    handle_head* head = (handle_head*)a0;
    QBoxLayout *self = (QBoxLayout*)head->native;
    switch (drvid) {
    case BOXLAYOUT_ADDLAYOUTWITH: {
        self->addLayout((QLayout*)drvGetNative(a1),drvGetInt(a2));
        break;
    }
    case BOXLAYOUT_ADDWIDGETWITH: {
        self->addWidget((QWidget*)drvGetNative(a1),drvGetInt(a2),Qt::Alignment(drvGetInt(a3)));
        break;
    }
    case BOXLAYOUT_ADDSPACING: {
        self->addSpacing(drvGetInt(a1));
        break;
    }
    case BOXLAYOUT_ADDSTRETCH: {
        self->addStretch(drvGetInt(a1));
        break;
    }
    default:
        return 0;
    }
    return 1;
}

int drv_hboxlayout(int drvid, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6, void* a7, void* a8, void* a9)
{
    handle_head* head = (handle_head*)a0;
    QHBoxLayout *self = (QHBoxLayout*)head->native;
    switch (drvid) {
    case HBOXLAYOUT_INIT: {
        drvNewObj(a0,new QHBoxLayout);
        break;
    }
    default:
        return 0;
    }
    return 1;
}

int drv_vboxlayout(int drvid, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6, void* a7, void* a8, void* a9)
{
    handle_head* head = (handle_head*)a0;
    QVBoxLayout *self = (QVBoxLayout*)head->native;
    switch (drvid) {
    case VBOXLAYOUT_INIT: {
        drvNewObj(a0,new QVBoxLayout);
        break;
    }
    default:
        return 0;
    }
    return 1;
}

int drv_basebutton(int drvid, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6, void* a7, void* a8, void* a9)
{
    handle_head* head = (handle_head*)a0;
    QAbstractButton *self = (QAbstractButton*)head->native;
    switch (drvid) {
    case BASEBUTTON_SETTEXT: {
        self->setText(drvGetString(a1));
        break;
    }
    case BASEBUTTON_TEXT: {
        drvSetString(a1,self->text());
        break;
    }
    default:
        return 0;
    }
    return 1;
}

int drv_button(int drvid, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6, void* a7, void* a8, void* a9)
{
    handle_head* head = (handle_head*)a0;
    QPushButton *self = (QPushButton*)head->native;
    switch (drvid) {
    case BUTTON_INIT: {
        drvNewObj(a0,new QPushButton);
        break;
    }
    case BUTTON_SETFLAT: {
        self->setFlat(drvGetBool(a1));
        break;
    }
    case BUTTON_ISFLAT: {
        drvSetBool(a1,self->isFlat());
        break;
    }
    case BUTTON_SETDEFAULT: {
        self->setDefault(drvGetBool(a1));
        break;
    }
    case BUTTON_ISDEFAULT: {
        drvSetBool(a1,self->isDefault());
        break;
    }
    case BUTTON_SETMENU: {
        self->setMenu(drvGetMenu(a1));
        break;
    }
    case BUTTON_MENU: {
        drvSetHandle(a1,self->menu());
        break;
    }
    case BUTTON_ONCLICKED: {
        QObject::connect(self,SIGNAL(clicked()),drvNewSignal(self,a1,a2),SLOT(call()));
        break;
    }
    default:
        return 0;
    }
    return 1;
}

int drv_checkbox(int drvid, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6, void* a7, void* a8, void* a9)
{
    handle_head* head = (handle_head*)a0;
    QCheckBox *self = (QCheckBox*)head->native;
    switch (drvid) {
    case CHECKBOX_INIT: {
        drvNewObj(a0,new QCheckBox);
        break;
    }
    case CHECKBOX_SETCHECK: {
        self->setCheckState((Qt::CheckState)drvGetInt(a1));
        break;
    }
    case CHECKBOX_CHECK: {
        drvSetInt(a1,self->checkState());
        break;
    }
    case CHECKBOX_SETTRISTATE: {
        self->setTristate(drvGetBool(a1));
        break;
    }
    case CHECKBOX_ISTRISTATE: {
        drvSetBool(a1,self->isTristate());
        break;
    }
    case CHECKBOX_ONSTATECHANGED: {
        QObject::connect(self,SIGNAL(stateChanged(int)),drvNewSignal(self,a1,a2),SLOT(call(int)));
        break;
    }
    default:
        return 0;
    }
    return 1;
}

int drv_radio(int drvid, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6, void* a7, void* a8, void* a9)
{
    handle_head* head = (handle_head*)a0;
    QRadioButton *self = (QRadioButton*)head->native;
    switch (drvid) {
    case RADIO_INIT: {
        drvNewObj(a0,new QRadioButton);
        break;
    }
    case RADIO_ONCLICKED: {
        QObject::connect(self,SIGNAL(clicked()),drvNewSignal(self,a1,a2),SLOT(call()));
        break;
    }
    default:
        return 0;
    }
    return 1;
}

int drv_frame(int drvid, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6, void* a7, void* a8, void* a9)
{
    handle_head* head = (handle_head*)a0;
    QFrame *self = (QFrame*)head->native;
    switch (drvid) {
    case FRAME_INIT: {
        drvNewObj(a0,new QFrame);
        break;
    }
    case FRAME_SETFRAMESTYLE: {
        self->setFrameStyle(drvGetInt(a1));
        break;
    }
    case FRAME_FRAMESTYLE: {
        drvSetInt(a1,self->frameStyle());
        break;
    }
    case FRAME_SETFRAMERECT: {
        self->setFrameRect(drvGetRect(a1));
        break;
    }
    case FRAME_FRAMERECT: {
        drvSetRect(a1,self->frameRect());
        break;
    }
    default:
        return 0;
    }
    return 1;
}

int drv_label(int drvid, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6, void* a7, void* a8, void* a9)
{
    handle_head* head = (handle_head*)a0;
    QLabel *self = (QLabel*)head->native;
    switch (drvid) {
    case LABEL_INIT: {
        drvNewObj(a0,new QLabel);
        break;
    }
    case LABEL_SETTEXT: {
        self->setText(drvGetString(a1));
        break;
    }
    case LABEL_TEXT: {
        drvSetString(a1,self->text());
        break;
    }
    case LABEL_SETWORDWRAP: {
        self->setWordWrap(drvGetBool(a1));
        break;
    }
    case LABEL_WORDWRAP: {
        drvSetBool(a1,self->wordWrap());
        break;
    }
    case LABEL_SETTEXTFORMAT: {
        self->setTextFormat((Qt::TextFormat)drvGetInt(a1));
        break;
    }
    case LABEL_TEXTFORMAT: {
        drvSetInt(a1,self->textFormat());
        break;
    }
    case LABEL_ONLINKACTIVATED: {
        QObject::connect(self,SIGNAL(linkActivated(QString)),drvNewSignal(self,a1,a2),SLOT(call(QString)));
        break;
    }
    default:
        return 0;
    }
    return 1;
}

int drv_groupbox(int drvid, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6, void* a7, void* a8, void* a9)
{
    handle_head* head = (handle_head*)a0;
    QGroupBox *self = (QGroupBox*)head->native;
    switch (drvid) {
    case GROUPBOX_INIT: {
        drvNewObj(a0,new QGroupBox);
        break;
    }
    case GROUPBOX_SETTITLE: {
        self->setTitle(drvGetString(a1));
        break;
    }
    case GROUPBOX_TITLE: {
        drvSetString(a1,self->title());
        break;
    }
    default:
        return 0;
    }
    return 1;
}

int drv_dialog(int drvid, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6, void* a7, void* a8, void* a9)
{
    handle_head* head = (handle_head*)a0;
    QDialog *self = (QDialog*)head->native;
    switch (drvid) {
    case DIALOG_INIT: {
        drvNewObj(a0,new QDialog);
        break;
    }
    case DIALOG_SETMODAL: {
        self->setModal(drvGetBool(a1));
        break;
    }
    case DIALOG_ISMODAL: {
        drvSetBool(a1,self->isModal());
        break;
    }
    case DIALOG_SETRESULT: {
        self->setResult(drvGetInt(a1));
        break;
    }
    case DIALOG_RESULT: {
        drvSetInt(a1,self->result());
        break;
    }
    case DIALOG_EXEC: {
        drvSetInt(a1,self->exec());
        break;
    }
    case DIALOG_DONE: {
        self->done(drvGetInt(a1));
        break;
    }
    case DIALOG_ACCEPT: {
        self->accept();
        break;
    }
    case DIALOG_REJECT: {
        self->reject();
        break;
    }
    case DIALOG_ONACCEPTED: {
        QObject::connect(self,SIGNAL(acceped()),drvNewSignal(self,a1,a2),SLOT(call()));
        break;
    }
    case DIALOG_ONREJECTED: {
        QObject::connect(self,SIGNAL(rejected()),drvNewSignal(self,a1,a2),SLOT(call()));
        break;
    }
    default:
        return 0;
    }
    return 1;
}

int drv_combobox(int drvid, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6, void* a7, void* a8, void* a9)
{
    handle_head* head = (handle_head*)a0;
    QComboBox *self = (QComboBox*)head->native;
    switch (drvid) {
    case COMBOBOX_INIT: {
        drvNewObj(a0,new QComboBox);
        break;
    }
    case COMBOBOX_COUNT: {
        drvSetInt(a1,self->count());
        break;
    }
    case COMBOBOX_SETCURRENTINDEX: {
        self->setCurrentIndex(drvGetInt(a1));
        break;
    }
    case COMBOBOX_CURRENTINDEX: {
        drvSetInt(a1,self->currentIndex());
        break;
    }
    case COMBOBOX_CURRENTTEXT: {
        drvSetString(a1,self->currentText());
        break;
    }
    case COMBOBOX_SETEDITABLE: {
        self->setEditable(drvGetBool(a1));
        break;
    }
    case COMBOBOX_ISEDITABLE: {
        drvSetBool(a1,self->isEditable());
        break;
    }
    case COMBOBOX_SETMAXCOUNT: {
        self->setMaxCount(drvGetInt(a1));
        break;
    }
    case COMBOBOX_MAXCOUNT: {
        drvSetInt(a1,self->maxCount());
        break;
    }
    case COMBOBOX_SETMAXVISIBLEITEMS: {
        self->setMaxVisibleItems(drvGetInt(a1));
        break;
    }
    case COMBOBOX_MAXVISIBLEITEMS: {
        drvSetInt(a1,self->maxVisibleItems());
        break;
    }
    case COMBOBOX_SETMINIMUMCONTENTSLENGTH: {
        self->setMinimumContentsLength(drvGetInt(a1));
        break;
    }
    case COMBOBOX_MINIMUNCONTENTSLENGHT: {
        drvSetInt(a1,self->minimumContentsLength());
        break;
    }
    case COMBOBOX_ADDITEM: {
        self->addItem(drvGetString(a1));
        break;
    }
    case COMBOBOX_INSERTITEM: {
        self->insertItem(drvGetInt(a1),drvGetString(a2));
        break;
    }
    case COMBOBOX_REMOVEITEM: {
        self->removeItem(drvGetInt(a1));
        break;
    }
    case COMBOBOX_ITEMTEXT: {
        drvSetString(a2,self->itemText(drvGetInt(a1)));
        break;
    }
    case COMBOBOX_ONCURRENTINDEXCHANGED: {
        QObject::connect(self,SIGNAL(currentIndexChanged(int)),drvNewSignal(self,a1,a2),SLOT(call(int)));
        break;
    }
    default:
        return 0;
    }
    return 1;
}

int drv_painter(int drvid, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6, void* a7, void* a8, void* a9)
{
    handle_head* head = (handle_head*)a0;
    QPainter *self = (QPainter*)head->native;
    switch (drvid) {
    case PAINTER_INIT: {
        drvNewHead(a0,new QPainter);
        break;
    }
    case PAINTER_DESTROY: {
        drvDelObj(a0,self);
        break;
    }
    case PAINTER_BEGIN: {
        self->begin(drvGetWidget(a1));
        break;
    }
    case PAINTER_END: {
        self->end();
        break;
    }
    case PAINTER_SETFONT: {
        self->setFont(drvGetFont(a1));
        break;
    }
    case PAINTER_FONT: {
        drvSetFont(a1,self->font());
        break;
    }
    case PAINTER_DRAWPOINT: {
        self->drawPoint(drvGetPoint(a1));
        break;
    }
    case PAINTER_DRAWLINE: {
        self->drawLine(drvGetPoint(a1),drvGetPoint(a2));
        break;
    }
    case PAINTER_DRAWLINES: {
        self->drawLines(drvGetPoints(a1));
        break;
    }
    case PAINTER_DRAWPOLYLINE: {
        self->drawPolyline(drvGetPoints(a1));
        break;
    }
    case PAINTER_DRAWTEXT: {
        self->drawText(drvGetPoint(a1),drvGetString(a2));
        break;
    }
    default:
        return 0;
    }
    return 1;
}

typedef int (*DRV_CALLBACK)(void* fn, void *a1,void* a2,void* a3,void* a4);

static DRV_CALLBACK pfn_drv_callback;

int drv_callback(void* fn, void *a1,void* a2,void* a3,void* a4)
{
    return pfn_drv_callback(fn,a1,a2,a3,a4);
}	

int drv(int drvcls, int drvid, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6, void* a7, void* a8, void* a9)
{
    switch(drvcls) {
    case -1:
        pfn_drv_callback = (DRV_CALLBACK)a0;
        return 1;
    case DRVCLASS_OBJECT:
        return drv_object(drvid,a0,a1,a2,a3,a4,a5,a6,a7,a8,a9);
    case DRVCLASS_APP:
        return drv_app(drvid,a0,a1,a2,a3,a4,a5,a6,a7,a8,a9);
    case DRVCLASS_WIDGET:
        return drv_widget(drvid,a0,a1,a2,a3,a4,a5,a6,a7,a8,a9);
    case DRVCLASS_MENUBAR:
        return drv_menubar(drvid,a0,a1,a2,a3,a4,a5,a6,a7,a8,a9);
    case DRVCLASS_MENU:
        return drv_menu(drvid,a0,a1,a2,a3,a4,a5,a6,a7,a8,a9);
    case DRVCLASS_ACTION:
        return drv_action(drvid,a0,a1,a2,a3,a4,a5,a6,a7,a8,a9);
    case DRVCLASS_TABWIDGET:
        return drv_tabwidget(drvid,a0,a1,a2,a3,a4,a5,a6,a7,a8,a9);
    case DRVCLASS_LAYOUT:
        return drv_layout(drvid,a0,a1,a2,a3,a4,a5,a6,a7,a8,a9);
    case DRVCLASS_BOXLAYOUT:
        return drv_boxlayout(drvid,a0,a1,a2,a3,a4,a5,a6,a7,a8,a9);
    case DRVCLASS_HBOXLAYOUT:
        return drv_hboxlayout(drvid,a0,a1,a2,a3,a4,a5,a6,a7,a8,a9);
    case DRVCLASS_VBOXLAYOUT:
        return drv_vboxlayout(drvid,a0,a1,a2,a3,a4,a5,a6,a7,a8,a9);
    case DRVCLASS_BASEBUTTON:
        return drv_basebutton(drvid,a0,a1,a2,a3,a4,a5,a6,a7,a8,a9);
    case DRVCLASS_BUTTON:
        return drv_button(drvid,a0,a1,a2,a3,a4,a5,a6,a7,a8,a9);
    case DRVCLASS_CHECKBOX:
        return drv_checkbox(drvid,a0,a1,a2,a3,a4,a5,a6,a7,a8,a9);
    case DRVCLASS_RADIO:
        return drv_radio(drvid,a0,a1,a2,a3,a4,a5,a6,a7,a8,a9);
    case DRVCLASS_FRAME:
        return drv_frame(drvid,a0,a1,a2,a3,a4,a5,a6,a7,a8,a9);
    case DRVCLASS_LABEL:
        return drv_label(drvid,a0,a1,a2,a3,a4,a5,a6,a7,a8,a9);
    case DRVCLASS_GROUPBOX:
        return drv_groupbox(drvid,a0,a1,a2,a3,a4,a5,a6,a7,a8,a9);
    case DRVCLASS_DIALOG:
        return drv_dialog(drvid,a0,a1,a2,a3,a4,a5,a6,a7,a8,a9);
    case DRVCLASS_COMBOBOX:
        return drv_combobox(drvid,a0,a1,a2,a3,a4,a5,a6,a7,a8,a9);
    case DRVCLASS_PAINTER:
        return drv_painter(drvid,a0,a1,a2,a3,a4,a5,a6,a7,a8,a9);
    default:
        return 0;
    }
    return 1;
}

