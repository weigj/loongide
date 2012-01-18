#include "qtdrv_global.h"
#include "qtsignal.h"
#include "qtevent.h"
#include "qtapp.h"

#include <QApplication>
#include <QWidget>
#include <QLayout>
#include <QMenu>
#include <QMetaType>
#include <QVariant>
#include <QEvent>
#include <QDebug>

enum {
    FUNC_TYPE_NONE = 0,
    FUNC_TYPE,
    FUNC_TYPE_S,
    FUNC_TYPE_RESIZE,
    FUNC_TYPE_LAST
};

typedef void* handle_ptr;

struct handle_head {
    void *native;
    int   classid;
};

struct string_head {
    const char *data;
    int len;
};

struct slice_head {
    const char *data;
    int len;
    int cap;
};

struct point_head {
    int x;
    int y;
};

struct size_head {
    int width;
    int hegiht;
};

struct rect_head {
    int x;
    int y;
    int width;
    int height;
};

struct margins_head {
    int left;
    int top;
    int right;
    int bottom;
};

typedef unsigned char byte;

struct drv_head_ptr
{
    drv_head_ptr(handle_head *p = 0) : ptr(p){}
    handle_head *ptr;
};
Q_DECLARE_METATYPE(drv_head_ptr)

struct drv_event_ptr
{
    drv_event_ptr() : ptr(0) {}
    QtEvent *ptr;
};
Q_DECLARE_METATYPE(drv_event_ptr)

inline QString drvGetString(void *param)
{
    if (param == 0) {
        return QString();
    }
    string_head *h = (string_head*)param;
    return QString::fromUtf8(h->data,h->len);
}

inline void drvSetString(void *param, const QString &s)
{
    if (param == 0) {
        return;
    }
    const QByteArray & ar = s.toUtf8();
    string_head *sh = (string_head*)param;
    sh->data = ar.constData();
    sh->len = ar.length();
}

inline QPoint drvGetPoint(void *param)
{
    if (param == 0) {
        return QPoint();
    }
    point_head *h = (point_head*)param;
    return QPoint(h->x,h->y);
}

inline QVector<QPoint> drvGetPoints(void *param)
{
    if (param == 0) {
        return QVector<QPoint> ();
    }
    slice_head *sh = (slice_head*)param;
    QVector<QPoint> vec(sh->len);
    point_head *h = (point_head*)sh->data;
    for (int i=0; i < sh->len; i++) {
        vec[i] = QPoint(h[i].x,h[i].y);
    }
    return vec;
}

inline void drvSetPoint(void *param, const QPoint &pt)
{
    if (param == 0) {
        return;
    }
    point_head *sh = (point_head*)param;
    sh->x = pt.x();
    sh->y = pt.y();
}

inline QSize drvGetSize(void *param)
{
    if (param == 0) {
        return QSize();
    }
    size_head *h = (size_head*)param;
    return QSize(h->width,h->hegiht);
}

inline void drvSetSize(void *param, const QSize &sz)
{
    if (param == 0) {
        return;
    }
    size_head *h = (size_head*)param;
    h->width = sz.width();
    h->hegiht = sz.height();
}

inline QRect drvGetRect(void *param)
{
    if (param == 0) {
        return QRect();
    }
    rect_head *h = (rect_head*)param;
    return QRect(h->x,h->y,h->width,h->height);
}

inline void drvSetRect(void *param, const QRect &rc)
{
    if (param == 0) {
        return;
    }
    rect_head *h = (rect_head*)param;
    h->x = rc.x();
    h->y = rc.y();
    h->width = rc.width();
    h->height = rc.height();
}

inline QMargins drvGetMargins(void *param)
{
    if (param == 0) {
        return QMargins();
    }
    margins_head *h = (margins_head*)param;
    return QMargins(h->left,h->top,h->right,h->bottom);
}

inline void drvSetMargins(void *param, const QMargins &mr)
{
    if (param == 0) {
        return;
    }
    margins_head *h = (margins_head*)param;
    h->left = mr.left();
    h->top = mr.top();
    h->right = mr.right();
    h->bottom = mr.bottom();
}

inline bool drvGetBool(void *param)
{
    return *(int*)param != 0;
}

inline void drvSetBool(void *param, bool b)
{
    *(int*)param = b?1:0;
}

inline int drvGetInt(void *param)
{
    return *(int*)param;
}

inline void drvSetInt(void *param, int value)
{
    *(int*)param = value;
}

typedef unsigned char byte;

struct font_head {
    string_head* family;
    int pointsize;
    int weight;
    int stretch;
    byte bold;
    byte italic;
    byte strikeout;
    byte underline;
    byte overline;
    byte fixedpitch;
};

inline QFont drvGetFont(void *param)
{
    if (param == 0) {
        return QFont();
    }
    font_head *h = (font_head*)param;
    QFont font;
    QString family = drvGetString(h->family);
    if (!family.isEmpty())
        font.setFamily(family);
    if (h->pointsize > 0)
        font.setPointSize(h->pointsize);
    font.setBold(font.Bold != 0);
    if (h->weight >= 0 && h->weight < 100) {
        font.setWeight(h->weight);
    }
    if (h->stretch > 0) {
        font.setStretch(h->stretch);
    }
    font.setItalic(h->italic != 0);
    font.setStrikeOut(h->strikeout != 0);
    font.setUnderline(h->underline != 0);
    font.setOverline(h->overline != 0);
    font.setFixedPitch(h->fixedpitch != 0);
    return font;
}

inline void drvSetFont(void *param, const QFont &font)
{
    if (param == 0) {
        return;
    }
    font_head *h = (font_head*)param;
    string_head sh;
    drvSetString(&sh,font.family());
    h->family = &sh;
    h->pointsize = font.pointSize();
    h->weight = font.weight();
    h->stretch = font.stretch();
    h->bold = font.bold();
    h->italic = font.italic();
    h->strikeout = font.strikeOut();
    h->underline = font.underline();
    h->overline = font.overline();
    h->fixedpitch = font.fixedPitch();
}

inline void drvSetHandle(void *param, QObject *obj)
{
    if (obj) {
        QVariant var = obj->property("qdrv_head");
        if (!var.isNull()) {
            handle_head *head = (handle_head*)param;
            *head = *var.value<drv_head_ptr>().ptr;
        }
    }
}

inline void* drvGetNative(void *param)
{
    if (param == 0) {
        return 0;
    }    
    return ((handle_head*)param)->native;
}

inline QWidget* drvGetWidget(void *param)
{
    if (param == 0) {
        return 0;
    }
    return (QWidget*)((handle_head*)param)->native;
}

inline QLayout* drvGetLayout(void *param)
{
    if (param == 0) {
        return 0;
    }
    return (QLayout*)((handle_head*)param)->native;
}

inline QMenu* drvGetMenu(void *param)
{
    if (param == 0) {
        return 0;
    }
    return (QMenu*)((handle_head*)param)->native;
}

inline QAction* drvGetAction(void *param)
{
    if (param == 0) {
        return 0;
    }
    return (QAction*)((handle_head*)param)->native;
}

inline void drvNewObj(void *a0, QObject *obj)
{
    handle_head *head =(handle_head*)a0;
    head->native = obj;
    obj->setProperty("qdrv_head",QVariant::fromValue(drv_head_ptr(head)));
    QObject::connect(obj,SIGNAL(destroyed(QObject*)),((QtApp*)qApp),SLOT(deleteObject(QObject*)));
}

inline void drvNewHead(void *a0, void *obj)
{
    handle_head *head =(handle_head*)a0;
    head->native = obj;
}

template <typename T>
inline void drvDelObj(void *a0, T *obj)
{
    if (a0 == 0 || obj == 0) {
        return;
    }
    handle_head *head =(handle_head*)a0;
    if (head->native == 0) {
        return;
    }
    delete obj;
}

inline QtSignal* drvNewSignal(QObject *parent, void *fn, void *param = 0)
{
    QtSignal *s = new QtSignal(0,fn);
    s->moveToThread(parent->thread());
    s->setParent(parent);
    return s;
}

inline void drvNewEvent(int type, void *a0, void *a1, void *a2 = 0)
{
    handle_head* head = (handle_head*)a0;
    QObject *obj = (QObject*)head->native;
    drv_event_ptr ptr;
    QVariant var = obj->property("drv_event");
    if (!var.isNull()) {
        ptr = var.value<drv_event_ptr>();
    } else {
        ptr.ptr = new QtEvent;
        ptr.ptr->moveToThread(obj->thread());
        ptr.ptr->setParent(obj);
        obj->installEventFilter(ptr.ptr);
    }
    ptr.ptr->setEvent(type,a1);
}

int drv_callback(void* fn, void *a1,void* a2,void* a3,void* a4);

extern "C"
int QTDRVSHARED_EXPORT drv(int classid, int drvid, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6, void* a7, void* a8, void* a9);
